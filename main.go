package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	. "twc/channel"
	. "twc/channel/channels"
	"twc/setup"
	"twc/ui"
	"twc/utils"
	. "twc/video"

	"golang.org/x/term"
)

func main() {
	setup.CreateProgramFiles()
	setup.LoadConfig()

	// flags
	openChannel := flag.String("o", "", "open channel")
	viewVods := flag.String("v", "", "view vods of channel")
	editChannelsFile := flag.Bool("e", false, "open channels file in default text editor")
	flag.Parse()

	switch {
	case *openChannel != "":
		handleOpenChannel(*openChannel)
		os.Exit(0)
	case *viewVods != "":
		handleViewVods(*viewVods)
		os.Exit(0)
	case *editChannelsFile:
		handleEditChannelsFile()
		os.Exit(0)
	}

	// default behavior

	// open script in new terminal if it was invoked outside of a tty (via gui launcher eg. dmenu)
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		exec.Command(os.Getenv("TERMINAL"), os.Args...).Run()
		os.Exit(0)
	}

	var channels Channels
	channels.GetChannels()
	channels.CheckStatus()
	channels.SortChannels()

	choice := ui.Menu(channels)

	switch choice := choice.(type) {
	case ui.ItemWrapper[Channel]:
		choice.Data.OpenChannel()
	case ui.ItemWrapper[Video]:
		choice.Data.Open()
	}
}

func handleOpenChannel(openChannel string) {
	channelParts := strings.Split(openChannel, " ")

	if len(channelParts) == 2 {
		channel := GetChannelType(channelParts[1])
		channel.SetName(channelParts[0])
		channel.OpenChannel()
	} else {
		var channels Channels
		channels.GetChannels()
		channels.FilterChannels(channelParts[0])
		if len(channels) != 0 {
			channels[0].OpenChannel()
		} else {
			fmt.Println("channel not found")
		}
	}
}

func handleViewVods(viewVods string) {
	channelParts := strings.Split(viewVods, " ")

	if len(channelParts) == 2 {
		channel := GetChannelType(channelParts[1])
		channel.SetName(channelParts[0])

		videos := channel.GetVods()
		if len(videos) == 0 {
			fmt.Println("channel does not have streams")
		} else {
			if video := ui.Menu(videos); video != nil {
				video.(ui.ItemWrapper[Video]).Data.Open()
			}
		}
	} else {
		var channels Channels
		channels.GetChannels()
		channels.FilterChannels(channelParts[0])

		if len(channels) == 0 {
			fmt.Println("channel not found")
		}

		videos := channels[0].GetVods()
		if len(videos) == 0 {
			fmt.Println("channel does not have streams")
		} else {
			if video := ui.Menu(videos); video != nil {
				video.(ui.ItemWrapper[Video]).Data.Open()
			}
		}
	}
}

func handleEditChannelsFile() {
	if channels_file := os.Getenv("CHANNELS_PATH"); channels_file != "" {
		utils.OpenWithDefaultApp(channels_file)
	}
}
