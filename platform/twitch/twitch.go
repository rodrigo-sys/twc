package platform

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"twc/types"
)

type Twitch struct{}

func (t Twitch) GetUrl(channel types.Channel) string {
	var url string
	url = fmt.Sprintf("https://www.twitch.tv/%s", channel.Name)

	return url
}

func (t Twitch) CheckStatus(channel types.Channel) bool {
	response, error := http.Get(channel.Platform.GetUrl(channel))
	if error != nil {
		log.Fatalln(error)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	return strings.Contains(string(body), "live_user")
	// return strings.Contains(string(body), "hqdefault_live.jpg")
}

func (t Twitch) OpenChannel(channel types.Channel) {
	url := channel.Platform.GetUrl(channel)
	exec.Command("mpv", url).Start()
	exec.Command("chatterino", "-c", channel.Name).Start()

	/*
		// sometimes mpv delays opening the video
		// so i watch for errors with tsp
		// or i try to use the direct url
		// td: think to use this as fallback

		// exec.Command("sh", "-c", fmt.Sprintf("tsp mpv %s", url)).Start()

		data, _ := exec.Command("sh", "-c", "yt-dlp --get-title --get-url "+url).Output()
		split_data = strings.Split(string(data), "\n")
		title := split_data[0] ; playback_url := split_data[1]
		exec.Command("sh", "-c",
			fmt.Sprintf(`mpv --title='%s' --force-media-title='%s' %s`, title, title, playback_url)).Start()
	*/
}

func (t Twitch) GetVods(channel types.Channel) []types.Video {
	url := channel.Platform.GetUrl(channel)
	videos_url := url + "/videos?filter=archives&sort=time"

	var videos []types.Video

	output, _ := exec.Command("sh", "-c",
		fmt.Sprintf(`yt-dlp --flat-playlist --lazy-playlist --playlist-items ':20' --print '%%(title)s' --print '%%(webpage_url)s' '%s'`, videos_url)).Output()

	fmt.Println(string(output))

	items := strings.Split(string(output), "\n")

	for i := 0; i < len(items)-1; i += 2 {
		videos = append(videos, types.Video{
			Name:     items[i],
			Url:      items[i+1],
			Platform: Twitch{},
		})
	}

	return videos
}

func (t Twitch) OpenVod(video types.Video) {
	exec.Command("mpv", video.Url).Start()
}
