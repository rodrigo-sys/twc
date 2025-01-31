package utils

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"syscall"
)

func CloudScraperGet(url string) string {
	output, _ := exec.Command("stealth-cli", url).Output()
	return string(output)
}

/*
func CloudScraperGet(url string) string {
	cmd := `import cloudscraper
scraper = cloudscraper.create_scraper()
print(scraper.get("%s").text)`

	output, _ := exec.Command("python", "-c", fmt.Sprintf(cmd, url)).Output()

	for {
		if output[0] != '<' {
			break
		}
		output, _ = exec.Command("python", "-c", fmt.Sprintf(cmd, url)).Output()
	}

	return string(output)
}
*/

func OpenWithDefaultApp(filePath string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", filePath)
	case "darwin":
		cmd = exec.Command("open", filePath)
	case "linux":
		cmd = exec.Command("xdg-open", filePath)

		desktop, _ := exec.Command("xdg-mime", "query", "default", "text/plain").Output()
		desktop_file, _ := os.ReadFile("/usr/share/applications/" + strings.TrimSpace(string(desktop)))
		if strings.Contains(string(desktop_file), "Terminal=true") {
			if os.Getenv("TERMINAL") == "" {
				break
			}

			for _, line := range strings.Split(string(desktop_file), "\n") {
				if strings.HasPrefix(line, "Exec=") {
					exec_value := strings.TrimSpace(strings.TrimPrefix(line, "Exec="))

					r := regexp.MustCompile("%f|%F")
					cmd_string := r.ReplaceAllString(exec_value, filePath)

					splited := strings.Split(cmd_string, " ")
					cmd = exec.Command(os.Getenv("TERMINAL"), splited...)
					break
				}
			}
		}

	default:
		return fmt.Errorf("unsupported os: %s", runtime.GOOS)
	}

	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	return cmd.Start()

	/*
		fmt.Printf("%#v\n", cmd.Args)
		output, err := cmd.CombinedOutput()
		fmt.Println(err.Error())
		fmt.Println(string(output))
		return nil
	*/
}

func ParseChatEnvar(chat_envar string, name string, url string) (exec.Cmd, error) {
	var cmd exec.Cmd
	var chat_string string

	if os.Getenv(chat_envar) == "" {
		return exec.Cmd{}, fmt.Errorf("chat envar empty")
	}
	chat_string = os.Getenv(chat_envar)

	if match, _ := regexp.MatchString(`^['"]`, chat_string); match {
		r, _ := regexp.Compile(`['"](.*)['"] (.*)`)
		match := r.FindStringSubmatch(chat_string)

		if len(match) != 0 {
			cmd.Args = append(cmd.Args, match[1])
			chat_string = match[2]
		}
	}

	cmd.Args = append(cmd.Args, strings.Split(chat_string, " ")...)
	path, _ := exec.LookPath(cmd.Args[0])
	cmd.Path = path

	// replace placeholder with actual values
	name_regex := regexp.MustCompile(`%n`)
	url_regex := regexp.MustCompile(`%u`)
	for i, arg := range cmd.Args {
		arg = name_regex.ReplaceAllString(arg, name)
		arg = url_regex.ReplaceAllString(arg, url)
		cmd.Args[i] = arg
	}

	return cmd, nil
}
