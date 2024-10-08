package main

import (
	. "twc/channels"
	config "twc/config/utils"
	. "twc/videos"
)

func main() {
	config.SetupConfig()

	var channels Channels

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
}
