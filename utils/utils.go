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
		desktop, _ := exec.Command("xdg-mime", "query", "default", "text/plain").Output()
		desktop_file, _ := os.ReadFile("/usr/share/applications/" + strings.TrimSpace(string(desktop)))
		if strings.Contains(string(desktop_file), "Terminal=true") {
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
		} else {
			cmd = exec.Command("xdg-open", filePath)
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
