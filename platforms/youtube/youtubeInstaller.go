package platforms

import (
	"fmt"
	. "twc/installer/base-installer"
	. "twc/installer/ospackage"
	. "twc/installer/utils"
)

type YoutubeInstaller struct {
	BaseInstaller
}

func (y YoutubeInstaller) InstallOfficialChat() error {
	// exit if pytchatty is already installed
	if IsInstalled("pytchatty") {
		return nil
	}

	// install git
	fmt.Println("installing git")
	err := InstallGit()
	if err != nil {
		return fmt.Errorf("Error intalling git: %w", err)
	}

	// install pytchatty
	fmt.Println("installing pytchatty")
	Ospackage{Name: "git+https://github.com/rodrigo-sys/pytchatty", Install_method: "pipx"}.Install()

	return nil
}
