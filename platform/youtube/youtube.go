package platform

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"twc/types"
)

type Youtube struct{}

func (y Youtube) GetUrl(channel types.Channel) string {
	var url string
	url = fmt.Sprintf("https://www.youtube.com/@%s/live", channel.Name)
	return url
}

func (y Youtube) CheckStatus(channel types.Channel) bool {
	response, error := http.Get(channel.Platform.GetUrl(channel))
	if error != nil {
		log.Fatalln(error)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	// return strings.Contains(string(body), "hqdefault_live.jpg")
	return strings.Contains(string(body), "isLiveDvrEnabled")
}

func (y Youtube) OpenChannel(channel types.Channel) {
	/* td: create openChat and openVideo fn and make this fn call them */

	url := channel.Platform.GetUrl(channel)
	webpage_url, _ := exec.Command("sh", "-c", fmt.Sprintf("yt-dlp --print webpage_url '%s'", url)).Output()

	pytchat_path := os.Getenv("PYTCHAT_PATH")
	exec.Command("mpv", url).Start()
	exec.Command("sh", "-c", fmt.Sprintf(`"$TERMINAL" sh -c "'%s' '%s'"`, pytchat_path, string(webpage_url))).Start()
}

func (y Youtube) GetVods(channel types.Channel) []types.Video {
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

	url := channel.Platform.GetUrl(channel)
	// videos_url := strings.Replace(url, "live", "videos", -1)
	videos_url := strings.Replace(url, "live", "streams", -1)

	var videos []types.Video

	output, _ := exec.Command("sh", "-c",
		fmt.Sprintf(`yt-dlp --extractor-args 'youtube:skip=hls,dash,translated_subs' --flat-playlist --lazy-playlist --playlist-items ':20' --print '%%(title)s' --print '%%(webpage_url)s' '%s'`, videos_url)).Output()

	items := strings.Split(string(output), "\n")

	for i := 0; i < len(items)-1; i += 2 {
		videos = append(videos, types.Video{
			Name:     items[i],
			Url:      items[i+1],
			Platform: Youtube{},
		})
	}

	return videos
}

func (y Youtube) OpenVod(video types.Video) {
	// td: handle and notify yt-dlp errors

	exec.Command("mpv", video.Url).Start()

	/*
		// using direct url
		url := video.Url
		data, _ := exec.Command("sh", "-c", "yt-dlp --get-title --get-url "+url).Output()
		split_data := strings.Split(string(data), "\n")

		title := split_data[0]
		playback_url := split_data[1]

		exec.Command("sh", "-c",
			fmt.Sprintf(`mpv --title='%s' --force-media-title='%s' %s`, title, title, playback_url)).Start()
	*/
}
