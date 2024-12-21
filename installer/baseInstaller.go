package installer
import (
	"fmt"
)
var (
	ErrorOfficialChatEmpty   = fmt.Errorf("official_chat cannot be empty")
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
