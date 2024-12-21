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
