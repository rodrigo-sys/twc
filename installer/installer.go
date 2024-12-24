package installer

import (
	"fmt"
	. "twc/installer/ospackage"
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

	// if chat is popout just exit
	if i.Custom_chat == "popout" {
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

/*
func (b BaseInstaller) InstallOfficialChat() error {
	fmt.Println("base InstallOfficialChat")
	return nil
}
*/

/*
func (b BaseInstaller) InstallOfficialPlayer() error {
	if b.Official_player != "mpv" {
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
*/
