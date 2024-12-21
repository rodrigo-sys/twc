package installer
import (
	"os/exec"
	"runtime"
)
type Ospackage struct {
	Name            string
	Install_method  string
	Aurhelper       string
func (p Ospackage) Install() error {
	if IsInstalled(p.Name) {
		return nil
	}

	switch p.Install_method {
	case "winget":
		return exec.Command("winget", "install", "--id", p.Name, "-e", "--accept-package-agreements", "--accept-source-agreements").Run()
	case "apt":
		return exec.Command("sh", "-c", "sudo apt-get install "+p.Name+" -y").Run()
	case "pacman":
		return exec.Command("sh", "-c", "sudo pacman -S "+p.Name+" --noconfirm").Run()
	case "aurhelper":
		if err := (Ospackage{Name: p.Aurhelper, Install_method: "pacman"}).Install(); err != nil {
			return err
		}
		return exec.Command("sh", "-c", p.Aurhelper+" -S "+p.Name+" --noconfirm").Run()
	case "pip":
		if err := InstallPip(); err != nil {
			return err
		}
		return exec.Command("python", "-m", "pip", "install", "--user", p.Name).Run()
	case "pipx":
		if err := InstallPipx(); err != nil {
			return err
		}
		return exec.Command("pipx", "install", p.Name).Run()
	}

	return nil
}
type system_type string
type Ospackages map[system_type]Ospackage
func (system_ospackages Ospackages) Install() error {
	ospackage, err := system_ospackages.Select()
	if err != nil {
		return err
	}

	return ospackage.Install()
}
func (system_ospackages Ospackages) Select() (Ospackage, error) {
	if pkg, ok := system_ospackages["any"]; ok {
		return pkg, nil
	}

	switch runtime.GOOS {
	case "linux":
		if pkg, ok := system_ospackages["linux"]; ok {
			return pkg, nil
		}

		if IsInstalled("apt") {
			if pkg, ok := system_ospackages["has_apt"]; ok {
				return pkg, nil
			}
		} else if IsInstalled("pacman") {
			if pkg, ok := system_ospackages["has_aurhelper"]; ok {
				if IsInstalled(pkg.Aurhelper) {
					return pkg, nil
				}
			}

			if pkg, ok := system_ospackages["has_pacman"]; ok {
				return pkg, nil
			}

		} else {
			return Ospackage{}, ErrUnsupportedPkgMgr
		}
	case "windows":
		if IsInstalled("winget") {
			if pkg, ok := system_ospackages["has_winget"]; ok {
				return pkg, nil
			}
		} else {
			return Ospackage{}, ErrWingetNotInstalled
		}
	}

	return Ospackage{}, nil
}
