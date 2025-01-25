package filescontent

import "strings"

var config = `
HOME="<your user home dir>"
CONFIG_DIR="<your config dir>"

TWC_CHANNELS_PATH="$CONFIG_DIR/twc/channels"

# You can use custom chats per platform like this

# KICKCHAT_PATH="$HOME/.npm-global/bin/kichatty"
# https://github.com/rodrigo-sys/kichatty

# YOUTUBECHAT_PATH="pytchatty"
# https://github.com/rodrigo-sys/pytchatty
`

var channels = `
#syntax:
#<username> <spaces or tab> <platform>

#examples:
#xQc	twitch
#Markiplier	youtube
#Trainwreckstv	kick
#Ninja	twitch
#Jacksepticeye	youtube
#Sodapoppin	kick
#Asmongold	twitch
#CorpseHusband	youtube
#Ludwig	twitch

#getting the username:

#you can get the username from the url
#https://www.twitch.tv/xqc
#https://kick.com/trainwreckstv

#in youtube, the username may not appear in the url 
#but can be found in below the channel name, prefixed with an @ symbol.
#https://www.youtube.com/channel/UC7_YxT-KID8kRbqZo7MyscQ
`

var Channels = strings.TrimSpace(channels)
var Config = strings.TrimSpace(config)
