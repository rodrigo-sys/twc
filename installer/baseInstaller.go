package installer
import (
	"fmt"
)
var (
	ErrorOfficialChatEmpty   = fmt.Errorf("official_chat cannot be empty")
	ErrorOfficialPlayerEmpty = fmt.Errorf("official_player cannot be empty")
)

type BaseInstaller struct {
	chat            string
	player          string
	official_chat   string
	official_player string
}
func (b BaseInstaller) InstallChat() error {
	if b.official_chat == "" {
		return ErrorOfficialChatEmpty
	}

	fmt.Println("installing chat")

	// if chat is popout just exit
	if b.chat == "popout" {
		return nil
	}

	if b.chat != "" && b.chat != b.official_chat {
		fmt.Println("installing " + b.chat)
		return DefaultInstallStrategy(b.chat)
	}

	return b.InstallOfficialChat()
}

func (b BaseInstaller) InstallPlayer() error {
	if b.official_player == "" {
		return ErrorOfficialPlayerEmpty
	}

	fmt.Println("installing player")
	if b.player != "" && b.player != b.official_player {
		fmt.Println("installing " + b.player)
		return DefaultInstallStrategy(b.player)
	}

	return b.InstallOfficialPlayer()
}

func (b BaseInstaller) InstallOfficialChat() error {
	return nil
}

func (b BaseInstaller) InstallOfficialPlayer() error {
	if b.chat != "mpv" {
		return nil
	}

	fmt.Println("installing ytdlp")
	if err := InstallYtdl(); err != nil {
		return fmt.Errorf("Error installing ytdlp: %w", err)
	}

	fmt.Println("installing mpv")
	if err := InstallMpv(); err != nil {
		return fmt.Errorf("Error installing mpv: %w", err)
	}

	return nil
}
