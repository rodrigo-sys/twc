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
	installers := map[string]Installer{
		"kick":    {Official_chat: "kichatty", Official_player: "mpv", OfficialInstallers: KickInstallers{}},
		"youtube": {Official_chat: "pytchatty", Official_player: "mpv", OfficialInstallers: YoutubeInstallers{}},
		"twitch":  {Official_chat: "chatterino", Official_player: "mpv", OfficialInstallers: TwitchInstallers{}},
	}

	// flags
	installFlagSet := flag.NewFlagSet("install", flag.ExitOnError)
	platforms := installFlagSet.String("platforms", "", "platforms to install")
	install_all := installFlagSet.Bool("install-all", false, "install all platforms")
	popout_all := installFlagSet.Bool("popout-all", false, "use popout as chat for all platforms")

	// custom chat and player flags
	for platform, _ := range installers {
		installFlagSet.String(platform+"-chat", "", "custom chat for "+platform)
		installFlagSet.String(platform+"-player", "", "custom player for "+platform)
	}

	// exit of here if not flags were setted
	if len(os.Args) < 3 {
		return
	}

	installFlagSet.Parse(os.Args[2:])

	// check for conflicting flags
	if *install_all && *platforms != "" {
		fmt.Println("Error: Cannot use both -platforms and -install-all flags at the same time.")
		os.Exit(1)
	}

	// set custom chat and player for each platform
	for platform, installer := range installers {
		if *popout_all {
			installer.Custom_chat = "popout"
		}

		chat := installFlagSet.Lookup(platform + "-chat")
		if chat != nil && chat.Value.String() != "" {
			installer.Custom_chat = chat.Value.String()
		}

		player := installFlagSet.Lookup(platform + "-player")
		if player != nil {
			installer.Custom_player = player.Value.String()
		}
		installers[platform] = installer
	}

	/*
		// loop installer map and print the chat and player for each platform
		for platform, installer := range installers {
			fmt.Printf("Platform: %s\n", platform)
			fmt.Printf("Chat: %s\n", installer.Custom_chat)
			// fmt.Printf("Player: %s\n", installer.Custom_player)
			// fmt.Printf("Official Chat: %s\n", installer.Official_chat)
			// fmt.Printf("Official Player: %s\n", installer.Official_player)
		}
	*/

	// define platforms_to_install based on the setted flags
	var platform_to_install []string
	if *install_all {
		platform_to_install = []string{"kick", "youtube", "twitch"}
	}
	if *platforms != "" {
		platform_to_install = strings.Split(*platforms, " ")
	}

	// run the installers
	for _, platform := range platform_to_install {
		if installer, exists := installers[platform]; exists {
			fmt.Println("installing " + platform)
			installer.Install()
		} else {
			fmt.Printf("Warning: No installer found for platform '%s'\n", platform)
		}
	}

	// exit after finishing if these flags were setted
	os.Exit(0)
}
