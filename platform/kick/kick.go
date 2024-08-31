package platform

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"twc/types"
	"twc/utils"
)

type Kick struct{}

func (k Kick) GetUrl(channel types.Channel) string {
	/* td: think whether return the api url or regular url */
	url := fmt.Sprintf("https://kick.com/api/v2/channels/%s", channel.Name)
	return url
}

func (k Kick) CheckStatus(channel types.Channel) bool {
	var scrapper_output string
	var scrapper_json map[string]interface{}
	var url = channel.Platform.GetUrl(channel) + "/livestream"

	scrapper_output = utils.CloudScraperGet(url)
	json.Unmarshal([]byte(scrapper_output), &scrapper_json)
	// fmt.Printf("%+v\n", scrapper_json)

	data := scrapper_json["data"]

	return data != nil
}

func (k Kick) OpenChannel(channel types.Channel) {
	url := "https://kick.com/" + channel.Name //+ "/livestream"
	exec.Command("mpv", url).Start()

	kickchat_dir := os.Getenv("KICKCHAT_DIR")
	exec.Command("sh", "-c", fmt.Sprintf(
		`"$TERMINAL" sh -c "(cd '%s' && pnpm dev '%s')"`, kickchat_dir, channel.Name,
	)).Start()

	/*
		// using direct url
		var scrapper_json map[string]interface{}
		var scrapper_output string

		scrapper_output = utils.CloudScraperGet(channel.Platform.GetUrl(channel) + "/livestream")
		json.Unmarshal([]byte(scrapper_output), &scrapper_json)

		playback_url := scrapper_json["data"].(map[string]interface{})["playback_url"].(string)
		title := scrapper_json["data"].(map[string]interface{})["session_title"].(string)

		exec.Command("sh", "-c", fmt.Sprintf(
			`mpv --title='%s' --force-media-title='%s' '%s'`, title, title, playback_url,
		)).Start()
	*/
}

func (k Kick) GetVods(channel types.Channel) []types.Video {
	/* td: think return Videos custom type instead */
	var videos []types.Video
	var scrapper_json []interface{}
	var scrapper_output string

	videos_endpoint := channel.Platform.GetUrl(channel) + "/videos"

	scrapper_output = utils.CloudScraperGet(videos_endpoint)
	json.Unmarshal([]byte(scrapper_output), &scrapper_json)

	//td: use https://kick.com/video/ + ["video"]["uuid"] instead of ["source"] as Url
	for i := range scrapper_json {
		videos = append(videos, types.Video{
			Url:      scrapper_json[i].(map[string]interface{})["source"].(string),
			Name:     scrapper_json[i].(map[string]interface{})["session_title"].(string),
			Platform: Kick{},
		})
	}

	/*
		pretty, _ := json.MarshalIndent(videos, "", "    ")
		fmt.Println(string(pretty))
	*/

	return videos
}

func (kick Kick) OpenVod(video types.Video) {
	exec.Command("mpv", video.Url).Start()
}
