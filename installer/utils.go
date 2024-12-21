package installer
import (
	"os"
	"os/exec"
	"path/filepath"
)
func IsInstalled(program string) bool {
	if !filepath.IsAbs(program) {
		_, err := exec.LookPath(program)
		return err == nil
	}

	_, err := os.Stat(program)
	return err == nil
	// return os.IsExist(err)
}
func DefaultInstallStrategy(ospackage_names ...string) error {
	var ospackages Ospackages
	if len(ospackage_names) == 3 {
		ospackages = Ospackages{
			"has_apt":    Ospackage{Name: ospackage_names[0], Install_method: "apt"},
			"has_pacman": Ospackage{Name: ospackage_names[1], Install_method: "pacman"},
			"has_winget": Ospackage{Name: ospackage_names[2], Install_method: "winget"},
		}
	} else {
		ospackages = Ospackages{
			"has_apt":    Ospackage{Name: ospackage_names[0], Install_method: "apt"},
			"has_pacman": Ospackage{Name: ospackage_names[0], Install_method: "pacman"},
			"has_winget": Ospackage{Name: ospackage_names[0], Install_method: "winget"},
		}
	}

	return ospackages.Install()
}
func InstallMpv() error {
	return DefaultInstallStrategy("mpv", "mpv", "mpv.net")
}

func InstallYtdl() error {
	return DefaultInstallStrategy("git", "git", "Git.Git")
}

func InstallGit() error {
	return DefaultInstallStrategy("git", "git", "Git.Git")
}
func InstallNpm() error {
	return DefaultInstallStrategy("npm", "npm", "OpenJS.NodeJS")
}
func InstallPip() error {
	return DefaultInstallStrategy("python3-pip", "python-pip", "python")
}
func InstallPipx() error {
	runEnsurePath := func() error {
		// Locate pipx script path
		localAppData := os.Getenv("LOCALAPPDATA")
		if localAppData == "" {
			return fmt.Errorf("LOCALAPPDATA not found.")
		}

		var pipxPath string
		rootPath := filepath.Join(localAppData, "Packages")
		_ = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
			if strings.Contains(path, "PythonSoftwareFoundation") && strings.Contains(path, "Scripts") {
				pipxPath = filepath.Join(path, "pipx.exe")
				return filepath.SkipDir
			}
			return nil
		})

		if pipxPath == "" {
			return fmt.Errorf("Pipx script path not found.")
		}

		// Run pipx ensurepath
		cmd := exec.Command(pipxPath, "ensurepath")
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("Error running pipx ensurepath:", err)
		}

		return nil
	}

	ospackages := Ospackages{
		"has_apt":    Ospackage{Name: "pipx", Install_method: "apt"},
		"has_pacman": Ospackage{Name: "python-pipx", Install_method: "pacman"},
		"has_winget": Ospackage{Name: "pipx", Install_method: "pip", Extra_setup: runEnsurePath},
	}

	return ospackages.Install()
}

func InstallFlatpak() error {
	Ospackages{
		"has_apt":    Ospackage{Name: "flatpak", Install_method: "apt"},
		"has_pacman": Ospackage{Name: "flatpak", Install_method: "pacman"},
	}.Install()

	return exec.Command("sh", "-c", "sudo flatpak remote-add --if-not-exists flathub https://dl.flathub.org/repo/flathub.flatpakrepo").Run()
}
