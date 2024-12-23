package installer

/*
This package can be improved, but for now, it is convenient.
*/

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var (
	ErrUnsupportedPkgMgr  = fmt.Errorf("not supported package manager")
	ErrWingetNotInstalled = fmt.Errorf("winget not installed")
	ErrRunningInstallCmd  = fmt.Errorf("error running install cmd")
	ErrPipxPath           = fmt.Errorf("error ensuring pipx in path")
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
