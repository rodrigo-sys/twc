package browser

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

type Browser struct {
	Cmd exec.Cmd
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
