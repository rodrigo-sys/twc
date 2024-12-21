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
