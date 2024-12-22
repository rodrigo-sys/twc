package platforms

import (
	"fmt"
	. "twc/installer/base-installer"
	. "twc/installer/ospackage"
	. "twc/installer/utils"
)

type TwitchInstaller struct {
	BaseInstaller
}

func (y TwitchInstaller) InstallOfficialChat() error {
	// exit if chatterino is already installed
	if IsInstalled("chatterino") || IsInstalled("com.chatterino.chatterino") || IsInstalled("/var/lib/flatpak/exports/bin/com.chatterino.chatterino") {
		return nil
	}

	// install chatterino
	fmt.Println("installing chatterino")
	return Ospackages{
		"has_apt":       Ospackage{Name: "app/com.chatterino.chatterino/x86_64/stable", Install_method: "flatpak"},
		"has_aurhelper": Ospackage{Name: "chatterino2-git ", Install_method: "aurhelper"},
		"has_pacman":    Ospackage{Name: "app/com.chatterino.chatterino/x86_64/stable", Install_method: "flatpak"},
		"has_winget":    Ospackage{Name: "ChatterinoTeam.Chatterino", Install_method: "winget"},
	}.Install()
}
