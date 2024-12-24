package installer

import (
	"flag"
	"fmt"
	"os"
	"strings"
	. "twc/installer"
	. "twc/platforms/kick"
	. "twc/platforms/twitch"
	. "twc/platforms/youtube"
)

func init() {
	// flags
	install := flag.String("install", "", "platforms to install")
	install_all := flag.Bool("install-all", false, "install all platforms")
	flag.Parse()

	// exit of here if not flags were setted
	if !(*install_all || *install != "") {
		return
	}

	// check for conflicting flags
	if *install_all && *install != "" {
		fmt.Println("Error: Cannot use both -install and -install-all flags at the same time.")
		os.Exit(1)
	}

	// define platforms_to_install based on the setted flags
	var platform_to_install []string
	if *install_all {
		platform_to_install = []string{"kick", "youtube", "twitch"}
	}
	if *install != "" {
		platform_to_install = strings.Split(*install, " ")
	}

	installers := map[string]Installer{
		"kick":    {Official_chat: "kichatty", Official_player: "mpv", OfficialInstallers: KickInstallers{}},
		"youtube": {Official_chat: "pytchatty", Official_player: "mpv", OfficialInstallers: YoutubeInstallers{}},
		"twitch":  {Official_chat: "chatterino", Official_player: "mpv", OfficialInstallers: TwitchInstallers{}},
	}

	for _, platform := range platform_to_install {
		if installer, exists := installers[platform]; exists {
			fmt.Println("installing " + platform)
			installer.Install()
		} else {
			// fmt.Println("platform " + platform + " not found")
			fmt.Printf("Warning: No installer found for platform '%s'\n", platform)
		}
	}

	// exit after finishing if these flags were setted
	if *install_all || *install != "" {
		os.Exit(0)
	}
}
