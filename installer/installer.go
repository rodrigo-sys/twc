package installer

import (
	"fmt"
	"os"
	config "twc/config/utils"
	. "twc/installer/ospackage"

	. "twc/platforms/kick"
	. "twc/platforms/twitch"
	. "twc/platforms/youtube"
)

type OfficialInstallers interface {
	InstallOfficialChat() error
	InstallOfficialPlayer() error
}

var (
	ErrorOfficialChatEmpty   = fmt.Errorf("official_chat cannot be empty")
	ErrorOfficialPlayerEmpty = fmt.Errorf("official_player cannot be empty")
)

type Installer struct {
	Custom_chat     string
	Custom_player   string
	Official_chat   string
	Official_player string
	OfficialInstallers
}

func (i Installer) Install() error {
	err_str := ""

	if err := i.InstallChat(); err != nil {
		err_str += fmt.Sprintf("%w", err)
	}

	if err := i.InstallPlayer(); err != nil {
		err_str += fmt.Sprintf("%w", err)
	}

	if err_str != "" {
		return fmt.Errorf(err_str)
	}

	return nil
}

func (i Installer) InstallChat() error {
	if i.Official_chat == "" {
		return ErrorOfficialChatEmpty
	}

	fmt.Println("installing chat")

	if i.Custom_chat == "popout" {
		i.setChatToPopoutInConfig()
		return nil
	}

	if i.Custom_chat != "" && i.Custom_chat != i.Official_chat {
		fmt.Println("installing " + i.Custom_chat)
		return DefaultInstallStrategy(i.Custom_chat)
	}

	return i.InstallOfficialChat()
}

func (i Installer) InstallPlayer() error {
	if i.Official_player == "" {
		return ErrorOfficialPlayerEmpty
	}

	fmt.Println("installing player")
	if i.Custom_player != "" && i.Custom_player != i.Official_player {
		fmt.Println("installing " + i.Custom_player)
		return DefaultInstallStrategy(i.Custom_player)
	}

	return i.InstallOfficialPlayer()
}

func (i Installer) setChatToPopoutInConfig() {
	var config_pair string
	switch i.OfficialInstallers.(type) {
	case KickInstallers:
		config_pair = "KICK_CHAT=popout"
	case TwitchInstallers:
		config_pair = "TWITCH_CHAT=popout"
	case YoutubeInstallers:
		config_pair = "YOUTUBE_CHAT=popout"
	}

	file, _ := os.OpenFile(config.GetConfigPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(config_pair)
}
