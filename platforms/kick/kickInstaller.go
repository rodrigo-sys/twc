package platforms

import (
	"fmt"
	. "twc/installer"
)

type kickInstaller struct {
	BaseInstaller
}

func (k kickInstaller) InstallOfficialChat() error {
	// exit if kichatty is already installed
	if IsInstalled("kichatty") || IsInstalled("kichatty") {
		return nil
	}

	// install kichatty
	fmt.Println("installing kichatty")
	return Ospackages{
		"any": Ospackage{Name: "https://github.com/rodrigo-sys/kichatty", Install_method: "npm"},
	}.Install()
}
