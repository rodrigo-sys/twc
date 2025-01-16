package platforms

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	. "twc/channel"
	. "twc/video"
)

type Youtube struct {
	BaseChannel
}

func (y Youtube) GetUrl() string {
	var url string
	url = fmt.Sprintf("https://www.youtube.com/@%s/live", y.BaseChannel.Name())
	return url
}

func (y Youtube) GetPopoutChatUrl() string {
	id, _ := exec.Command("yt-dlp", "--print", "id", y.GetUrl()).Output()
	return fmt.Sprintf("https://www.youtube.com/live_chat?is_popout=1&v=%s", id)
}

func (y Youtube) CheckStatus() bool {
	response, error := http.Get(y.GetUrl())
	if error != nil {
		log.Fatalln(error)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	// return strings.Contains(string(body), "hqdefault_live.jpg")
	return strings.Contains(string(body), "isLiveDvrEnabled")
}

func (y Youtube) OpenChannel() {
	/* td: create openChat and openVideo fn and make this fn call them */

	url := y.GetUrl()
	webpage_url, _ := exec.Command("sh", "-c", fmt.Sprintf("yt-dlp --print webpage_url '%s'", url)).Output()

	// open stream in player
	exec.Command("mpv", url).Start()

	// open chat
	if os.Getenv("YOUTUBECHAT_PATH") == "" {
		y.OpenPopoutChat(y.GetPopoutChatUrl())
	} else {
		exec.Command(os.Getenv("YOUTUBECHAT_PATH"), string(webpage_url)).Start()
	}
}

func (y Youtube) GetVods() Videos {
	/*
		td:
		analize which yt-dlp flags use.
		return Videos type.
		try use pure go or yt-dlp wrapers libraries.
		think a way to show videos and streams together
			get the 10 last videos and streams
			sort by timestamp
			make this configurable
	*/

	url := y.GetUrl()
	// videos_url := strings.Replace(url, "live", "videos", -1)
	videos_url := strings.Replace(url, "live", "streams", -1)

	var videos Videos

	output, _ := exec.Command("sh", "-c",
		fmt.Sprintf(`yt-dlp --extractor-args 'youtube:skip=hls,dash,translated_subs' --flat-playlist --lazy-playlist --playlist-items ':20' --print '%%(title)s' --print '%%(webpage_url)s' '%s'`, videos_url)).Output()

	items := strings.Split(string(output), "\n")

	for i := 0; i < len(items)-1; i += 2 {
		var video YoutubeVideo
		video.SetName(items[i])
		video.SetUrl(items[i+1])
		videos = append(videos, &video)
	}

	return videos
}

/* video */
type YoutubeVideo struct{ BaseVideo }

// func (y Youtube) OpenVod(video Video) {
// 	// td: handle and notify yt-dlp errors
//
// 	exec.Command("mpv", video.Url).Start()
//
// 	/*
// 		// using direct url
// 		url := video.Url
// 		data, _ := exec.Command("sh", "-c", "yt-dlp --get-title --get-url "+url).Output()
// 		split_data := strings.Split(string(data), "\n")
//
// 		title := split_data[0]
// 		playback_url := split_data[1]
//
// 		exec.Command("sh", "-c",
// 			fmt.Sprintf(`mpv --title='%s' --force-media-title='%s' %s`, title, title, playback_url)).Start()
// 	*/
// }
