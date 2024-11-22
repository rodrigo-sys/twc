package main

import (
	. "twc/channels"
	config "twc/config/utils"
	. "twc/types"
	"twc/ui"
	// . "twc/videos"
)

func main() {
	config.SetupConfig()

	var channels Channels
	channels.GetChannels()
	channels.CheckStatus()
	channels.SortChannels()

	choice := ui.Menu(channels)

	switch choice := choice.(type) {
	case ui.ItemWrapper[Channel]:
		choice.Data.Platform.OpenChannel(choice.Data)
	case ui.ItemWrapper[Video]:
		choice.Data.Platform.OpenVod(choice.Data)
	}

	/* legacy menu
	channels.GetChannels()
	channels.CheckStatus()
	channels.SortChannels()

	selected := channels.Menu()

	if selected.Islive {
		selected.Platform.OpenChannel(selected)
	} else {
		//td: return Videos type in GetVods
		var videos Videos
		videos = selected.Platform.GetVods(selected)
		vod := videos.Menu()
		vod.Platform.OpenVod(vod)
	}
	*/
}
