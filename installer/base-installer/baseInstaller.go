package installer

import (
	"fmt"
	. "twc/installer/ospackage"
)

type Installer interface {
	Install() error
	InstallChat() error
	InstallPlayer() error
	InstallOfficialChat() error
	InstallOfficialPlayer() error
}

var (
	ErrorOfficialChatEmpty   = fmt.Errorf("official_chat cannot be empty")
	ErrorOfficialPlayerEmpty = fmt.Errorf("official_player cannot be empty")
)

type BaseInstaller struct {
	Chat            string
	Player          string
	Official_chat   string
	Official_player string
}

func (b BaseInstaller) Install() error {
	err_str := ""

	if err := b.InstallChat(); err != nil {
		err_str += fmt.Sprintf("%w", err)
	}

	if err := b.InstallPlayer(); err != nil {
		err_str += fmt.Sprintf("%w", err)
	}

	if err_str != "" {
		return fmt.Errorf(err_str)
	}

	return nil
}

func (b BaseInstaller) InstallChat() error {
	if b.Official_chat == "" {
		return ErrorOfficialChatEmpty
	}

	fmt.Println("installing chat")

	// if chat is popout just exit
	if b.Chat == "popout" {
		return nil
	}

	if b.Chat != "" && b.Chat != b.Official_chat {
		fmt.Println("installing " + b.Chat)
		return DefaultInstallStrategy(b.Chat)
	}

	return b.InstallOfficialChat()
}

func (b BaseInstaller) InstallPlayer() error {
	if b.Official_player == "" {
		return ErrorOfficialPlayerEmpty
	}

	fmt.Println("installing player")
	if b.Player != "" && b.Player != b.Official_player {
		fmt.Println("installing " + b.Player)
		return DefaultInstallStrategy(b.Player)
	}

	return b.InstallOfficialPlayer()
}

func (b BaseInstaller) InstallOfficialChat() error {
	return nil
}

func (b BaseInstaller) InstallOfficialPlayer() error {
	if b.Chat != "mpv" {
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
