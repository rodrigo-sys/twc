package channel

import . "twc/video"

type Channel interface {
	CheckStatus() bool
	GetUrl() string
	GetVods() Videos
	OpenChannel()
	GetPopoutChatUrl() string
	OpenPopoutChat()

	Name() string
	Islive() bool
	Position() int

	SetName(name string)
	SetIslive(islive bool)
	SetPosition(position int)
}

type BaseChannel struct {
	name     string
	position int
	islive   bool
}

func (c *BaseChannel) GetUrl() string {
	return ""
}
func (c *BaseChannel) CheckStatus() bool {
	return false
}
func (c *BaseChannel) OpenChannel() {
}
func (c *BaseChannel) GetVods() Videos {
	return Videos{}
}

func (c *BaseChannel) GetPopoutChatUrl() string {
	return ""
}
func (c *BaseChannel) OpenPopoutChat(url string) {
}
/* getters */
func (c *BaseChannel) Name() string {
	return c.name
}
func (c *BaseChannel) Position() int {
	return c.position
}
func (c *BaseChannel) Islive() bool {
	return c.islive
}

/* setters */
func (c *BaseChannel) SetName(name string) {
	c.name = name
}
func (c *BaseChannel) SetIslive(islive bool) {
	c.islive = islive
}
func (c *BaseChannel) SetPosition(position int) {
	c.position = position
}
