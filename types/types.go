package types

/* td:
i want to put this custom types in different files
but that gives me an import cycle error

i will try to fix that later
*/

type Platform interface {
	CheckStatus(channel Channel) bool
	GetUrl(channel Channel) string
	GetVods(channel Channel) []Video
	OpenChannel(channel Channel)
	OpenVod(video Video)
}

// td: maybe change Name for Title
type Video struct {
	Name     string
	Url      string
	Platform Platform
}

type Channel struct {
	Name     string
	Platform Platform
	Position int
	Islive   bool
}
