package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	. "twc/channel"
	. "twc/channel/channels"
	config "twc/config/utils"
	"twc/ui"
	. "twc/video"
)

func main() {
	config.SetupConfig()

	// flags
	openChannel := flag.String("o", "", "open channel")
	flag.Parse()

	switch {
	case *openChannel != "":
		handleOpenChannel(*openChannel)
		os.Exit(0)
	}

	// default behavior
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
		if len(channels) == 0 {
			channels[0].OpenChannel()
		} else {
			fmt.Println("channel not found")
		}
	}
}
