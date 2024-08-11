package videos

import (
	"log"
	. "twc/types"

	"github.com/koki-develop/go-fzf"
)

type Videos []Video

func (videos Videos) Menu() Video {
	f, err := fzf.New()
	if err != nil {
		log.Fatal(err)
	}

	idxs, err := f.Find(videos, func(i int) string { return videos[i].Name })
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range idxs {
		return videos[i]
	}

	return Video{}
}
