package channels

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"

	. "twc/channel"
	// _ "twc/config"

	. "twc/platforms/kick"
	. "twc/platforms/twitch"
	. "twc/platforms/youtube"

	// . "twc/platform/youtube"

	// . "twc/video"

	"github.com/koki-develop/go-fzf"
)

type Channels []Channel

// td: this can be a function/method or it can be done another way
/*
var channels_types = map[string]IChannel{
	"twitch": &Twitch{},
	// "youtube": Youtube{},
	// "kick":    Kick{},
}
*/
func GetChannelType(type_string string) Channel {
	switch type_string {
	case "twitch":
		return &Twitch{}
	case "youtube":
		return &Youtube{}
	case "kick":
		return &Kick{}
	}
	return nil
}

func (channels *Channels) CheckStatusSync() {
	for i := range *channels {
		(*channels)[i].SetIslive((*channels)[i].CheckStatus())
		// (*channels)[i].Islive = (*channels)[i].Platform.CheckStatus((*channels)[i])
	}
}

func (channels *Channels) CheckStatus() {
	//td: make channel size configurable
	var wg sync.WaitGroup
	sem := make(chan struct{}, 14)

	for i := range *channels {
		sem <- struct{}{}
		wg.Add(1)

		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			(*channels)[i].SetIslive((*channels)[i].CheckStatus())
			// (*channels)[i].Islive = (*channels)[i].Platform.CheckStatus((*channels)[i])
		}()
	}

	wg.Wait()
}

func (channels *Channels) GetChannels() { // or followed
	// td: allow using official APIs to get the followed channels

	channels_file := os.Getenv("TWC_CHANNELS_PATH")
	content, err := os.ReadFile(channels_file)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	(*channels) = parseChannels(string(content))
}

func parseChannels(channels_raw string) Channels {
	// remove lines that start with #
	channels_raw = regexp.MustCompile("(?m)^#.*$").ReplaceAllString(channels_raw, "")
	channels_raw = regexp.MustCompile("\n\n+").ReplaceAllString(channels_raw, "\n")
	channels_raw = regexp.MustCompile(" +").ReplaceAllString(channels_raw, "\t")

	// create a new reader
	reader := csv.NewReader(strings.NewReader(channels_raw))
	reader.Comma = '\t' // Set the delimiter to tab

	// read all records from the file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	/* td: Islive field initial value must be neutral */
	var channels Channels
	for index, record := range records {
		channel := GetChannelType(record[1])

		channel.SetName(record[0])
		channel.SetPosition(index)
		channel.SetIslive(false)

		/*
			channel := Channel{
				Name:     record[0],
				Platform: platforms[record[1]],
				Position: index,
				Islive:   false,
			}
		*/

		channels = append(channels, channel)
	}

	return channels
}

func (channels *Channels) SortChannels() {
	// td: understand this, and remove the commented if
	sort.Slice((*channels), func(i int, j int) bool {
		// if channels[i].islive && channels[j].islive {
		if (*channels)[i].Islive() == (*channels)[j].Islive() {
			return (*channels)[i].Islive()
		}

		return !(*channels)[j].Islive()
	})
}

func (channels Channels) Menu() Channel {
	f, err := fzf.New()
	if err != nil {
		log.Fatal(err)
	}

	var idxs []int

	idxs, err = f.Find(channels, func(i int) string { return channels[i].Name() })

	if err != nil {
		log.Fatal(err)
	}

	for _, i := range idxs {
		return channels[i]
	}

	return &BaseChannel{}
}

func (channels *Channels) FilterChannels(name string) /*Channels*/ {
	/* td:
	think whether to return new array or modify the original.
	filter by any Channel struct field
	*/

	var filtered_channels Channels
	for i := range *channels {
		if (*channels)[i].Name() == name {
			filtered_channels = append(filtered_channels, (*channels)[i])
		}
	}

	(*channels) = filtered_channels
}
