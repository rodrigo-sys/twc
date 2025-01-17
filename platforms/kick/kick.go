package platforms

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	. "twc/channel"
	"twc/utils"
	. "twc/video"
)

type Kick struct {
	BaseChannel
}

func (k Kick) GetUrl() string {
	/* td: think whether return the api url or regular url */
	url := fmt.Sprintf("https://kick.com/api/v2/channels/%s", k.BaseChannel.Name())
	return url
}

func (k Kick) GetPopoutChatUrl() string {
	return fmt.Sprintf("https://kick.com/popout/%s/chat", k.BaseChannel.Name())
}

func (k Kick) CheckStatus() bool {
	/*
		err := exec.Command("sh", "-c", fmt.Sprintf("yt-dlp --print live_status --no-warnings '%s'", channel.Platform.GetUrl(channel))).Run()
		return err == nil
	*/

	// /*
	var scrapper_output string
	var scrapper_json map[string]interface{}
	var url = k.GetUrl() + "/livestream"

	scrapper_output = utils.CloudScraperGet(url)
	json.Unmarshal([]byte(scrapper_output), &scrapper_json)
	// fmt.Printf("%+v\n", scrapper_json)

	data := scrapper_json["data"]

	return data != nil
	// */
}

func (k Kick) OpenChannel() {
	url := "https://kick.com/" + k.BaseChannel.Name() //+ "/livestream"

	// open stream in player
	cmd := exec.Command("mpv", url)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true, Setctty: false}
	cmd.Start()

	// open chat
	if os.Getenv("KICKCHAT_PATH") == "" {
		k.OpenPopoutChat(k.GetPopoutChatUrl())
	} else {
		exec.Command(os.Getenv("KICKCHAT_PATH"), k.BaseChannel.Name()).Start()
	}
	/*
		// using direct url
		var scrapper_json map[string]interface{}
		var scrapper_output string

		scrapper_output = utils.CloudScraperGet(k.GetUrl() + "/livestream")
		json.Unmarshal([]byte(scrapper_output), &scrapper_json)

		playback_url := scrapper_json["data"].(map[string]interface{})["playback_url"].(string)
		title := scrapper_json["data"].(map[string]interface{})["session_title"].(string)

		exec.Command("sh", "-c", fmt.Sprintf(
			`mpv --title='%s' --force-media-title='%s' '%s'`, title, title, playback_url,
		)).Start()
	*/
}

func (k Kick) GetVods() Videos {
	/* td: think return Videos custom type instead */
	var videos Videos
	var scrapper_json []interface{}
	var scrapper_output string

	videos_endpoint := k.GetUrl() + "/videos"

	scrapper_output = utils.CloudScraperGet(videos_endpoint)
	json.Unmarshal([]byte(scrapper_output), &scrapper_json)

	//td: use https://kick.com/video/ + ["video"]["uuid"] instead of ["source"] as Url
	for i := range scrapper_json {
		var video KickVideo
		video.SetName(scrapper_json[i].(map[string]interface{})["session_title"].(string))
		video.SetUrl(scrapper_json[i].(map[string]interface{})["source"].(string))
		videos = append(videos, &video)
	}

	/*
		pretty, _ := json.MarshalIndent(videos, "", "    ")
		fmt.Println(string(pretty))
	*/

	return videos
}

/* video */
type KickVideo struct{ BaseVideo }
