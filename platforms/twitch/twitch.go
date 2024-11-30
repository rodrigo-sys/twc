package platforms

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
	. "twc/channel"
	. "twc/video"
)

/* channel */
type Twitch struct {
	BaseChannel
}

func (t Twitch) GetUrl() string {
	var url string
	url = fmt.Sprintf("https://www.twitch.tv/%s", t.BaseChannel.Name())

	return url
}

func (t Twitch) CheckStatus() bool {
	response, error := http.Get(t.GetUrl())
	if error != nil {
		log.Fatalln(error)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	return strings.Contains(string(body), "live_user")
	// return strings.Contains(string(body), "hqdefault_live.jpg")
}

func (t Twitch) OpenChannel() {
	url := t.GetUrl()
	exec.Command("mpv", url).Start()
	exec.Command("chatterino", "-c", t.BaseChannel.Name()).Start()

	/*
		// sometimes mpv delays opening the video
		// so i watch for errors with tsp
		// or i try to use the direct url
		// td: think to use this as fallback

		// exec.Command("sh", "-c", fmt.Sprintf("tsp mpv %s", url)).Start()

		data, _ := exec.Command("sh", "-c", "yt-dlp --get-title --get-url "+url).Output()
		split_data := strings.Split(string(data), "\n")
		title := split_data[0]
		playback_url := split_data[1]
		exec.Command("sh", "-c",
			fmt.Sprintf(`mpv --title='%s' --force-media-title='%s' %s`, title, title, playback_url)).Start()
	*/
}

func (t Twitch) GetVods() Videos {
	url := t.GetUrl()
	videos_url := url + "/videos?filter=archives&sort=time"

	var videos Videos

	output, _ := exec.Command("sh", "-c",
		fmt.Sprintf(`yt-dlp --flat-playlist --lazy-playlist --playlist-items ':20' --print '%%(title)s' --print '%%(webpage_url)s' '%s'`, videos_url)).Output()

	items := strings.Split(string(output), "\n")

	for i := 0; i < len(items)-1; i += 2 {
		var video TwitchVideo
		video.SetName(items[i])
		video.SetUrl(items[i+1])
		videos = append(videos, &video)
	}

	return videos
}

/* video */
type TwitchVideo struct{ BaseVideo }
