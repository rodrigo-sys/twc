package platforms

import (
	"fmt"
	. "twc/installer/ospackage"
	. "twc/installer/utils"
)

type KickInstallers struct {
}

func (k KickInstallers) InstallOfficialChat() error {
	// exit if kichatty is already installed
	if IsInstalled("kichatty") {
		return nil
	}

	// install kichatty
	fmt.Println("installing kichatty")
	return Ospackages{
		"any": Ospackage{Name: "https://github.com/rodrigo-sys/kichatty", Install_method: "npm"},
	}.Install()
}

func (k KickInstallers) InstallOfficialPlayer() error {
	fmt.Println("installing mpv")
	return InstallMpv()
}
