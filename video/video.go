package video

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"
	"twc/utils"

	"github.com/koki-develop/go-fzf"
)

type Videos []Video

type Video interface {
	Open()
	Name() string
	Url() string

	SetName(name string)
	SetUrl(url string)
}

// td: maybe change Name for Title
type BaseVideo struct {
	name string
	url  string
}

/* methods */
func (v *BaseVideo) Open() {
	cmd, _ := utils.ParsePlayerOption(os.Getenv("TWC_PLAYER"), v.url)

	// detach command
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	// create a pipe for the command's standard output
	stdout, _ := cmd.StdoutPipe()

	// start command
	cmd.Start()

	// show realtime output
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
}

/* getters */
func (v *BaseVideo) Name() string {
	return v.name
}
func (v *BaseVideo) Url() string {
	return v.url
}

/* setters */
func (v *BaseVideo) SetName(name string) {
	v.name = name
}
func (v *BaseVideo) SetUrl(url string) {
	v.url = url
}

/* this is the legacy menu */
func (videos Videos) Menu() Video {
	f, err := fzf.New()
	if err != nil {
		log.Fatal(err)
	}

	idxs, err := f.Find(videos, func(i int) string { return videos[i].Name() })
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range idxs {
		return videos[i]
	}

	return &BaseVideo{}
}
