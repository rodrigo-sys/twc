package main

import (
	. "twc/channel"
	. "twc/channel/channels"
	config "twc/config/utils"
	_ "twc/installer/process-flags"
	"twc/ui"
	. "twc/video"
)

func main() {
	config.SetupConfig()

	// /*
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
	// */

	/* legacy menu
	var channels Channels
	channels.GetChannels()
	channels.CheckStatus()
	channels.SortChannels()

	selected := channels.Menu()

	if selected.Islive() {
		selected.OpenChannel()
	} else {
		//td: return Videos type in GetVods
		var videos Videos
		videos = selected.GetVods()
		vod := videos.Menu()
		vod.Open()
	}
	*/
}
