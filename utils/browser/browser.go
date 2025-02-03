package browser

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"syscall"
	"time"
)

type Browser struct {
	Cmd exec.Cmd
}

func (b Browser) OpenUrl(url string) error {
	cmd := exec.Cmd{Path: b.Cmd.Path}

	r := regexp.MustCompile(`%u|%U`)
	if r.MatchString(strings.Join(b.Cmd.Args, " ")) {
		for _, arg := range b.Cmd.Args {
			cmd.Args = append(cmd.Args, r.ReplaceAllString(arg, url))
		}
	} else {
		cmd.Args = append(b.Cmd.Args, url)
	}

	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true, Setctty: false}
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	fmt.Println(output)
	// }
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start browser: %w", err)
	}

	// Run cmd.Wait() in a separate goroutine to avoid blocking
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait() // Send the result of Wait() to the channel
	}()

	select {
	case err := <-done:
		// Process finished, return the error (if any)
		if err != nil {
			return fmt.Errorf("browser exited with error: %w", err)
		}
	case <-time.After(1 * time.Second): // Allow time to detect immediate errors
		// Timeout reached, assume the browser is running fine
		// fmt.Printf("Browser launched with PID %d\n", cmd.Process.Pid)
		return nil
	}

	return nil
}

func NewBrowser() Browser {
	var cmd exec.Cmd
	var browser_string string

	if os.Getenv("TWC_BROWSER") == "" {
		browser_string, _ = GetDefaultBrowser()
	} else {
		browser_string = os.Getenv("TWC_BROWSER")
	}

	if match, _ := regexp.MatchString(`^['"]`, browser_string); match {
		r, _ := regexp.Compile(`['"](.*)['"]( (.*))?`)
		match := r.FindStringSubmatch(browser_string)

		if strings.HasPrefix(match[1], "~/") {
			r := regexp.MustCompile(`^~`)
			home, _ := os.UserHomeDir()
			match[1] = r.ReplaceAllString(match[1], home)
		}

		cmd.Args = append(cmd.Args, match[1])
		browser_string = match[3]
	}

	if browser_string != "" {
		cmd.Args = append(cmd.Args, strings.Split(browser_string, " ")...)
	}

	path, err := exec.LookPath(cmd.Args[0])
	if err != nil {
		fmt.Println(err.Error())
	}
	cmd.Path = path

	return Browser{Cmd: cmd}
}

func GetDefaultBrowser() (string, error) {
	// user BROWSER normal tamebien
	var browserPath string

	switch runtime.GOOS {
	case "windows":
		out, err := exec.Command("powershell", "-Command", "(Get-ItemProperty 'HKCU:\\Software\\Microsoft\\Windows\\Shell\\Associations\\UrlAssociations\\http').ProgId").Output()
		if err != nil {
			return "", fmt.Errorf("failed to get ProgID: %w", err)
		}
		progID := strings.TrimSpace(string(out))

		out, err = exec.Command("powershell", "-Command", fmt.Sprintf("(Get-ItemProperty 'HKCU:\\Software\\Clients\\StartMenuInternet\\%s').(default)", progID)).Output()
		if err != nil {
			return "", fmt.Errorf("failed to get browser executable: %w", err)
		}
		browserPath = strings.TrimSpace(string(out))

	case "darwin":
		out, err := exec.Command("osascript", "-e", "tell application \"System Events\" to get the POSIX path of (first item of (get url \"http://\"))").Output()
		if err != nil {
			return "", fmt.Errorf("failed to get default browser on macOS: %w", err)
		}
		browserPath = strings.TrimSpace(string(out))

	case "linux":
		out, err := exec.Command("xdg-settings", "get", "default-web-browser").Output()
		if err != nil {
			return "", fmt.Errorf("failed to get default browser on Linux: %w", err)
		}
		browser := strings.TrimSpace(string(out))

		desktopFilePath := fmt.Sprintf("/usr/share/applications/%s", browser)
		file, err := os.ReadFile(desktopFilePath)
		if err != nil {
			return "", fmt.Errorf("failed to read .desktop file: %w", err)
		}

		for _, line := range strings.Split(string(file), "\n") {
			if strings.HasPrefix(line, "Exec=") {
				browserPath = strings.TrimSpace(strings.TrimPrefix(line, "Exec="))
				break
			}
		}

	default:
		return "", fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	return browserPath, nil
}
