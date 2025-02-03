package filescontent

import "strings"

var config = `
HOME="<your user home dir>"
CONFIG_DIR="<your config dir>"

TWC_CHANNELS_PATH="$CONFIG_DIR/twc/channels"
TWC_PLAYER="mpv %u"

# You can use custom chats per platform like this
# placeholders: %n username, %u video/stream url

# KICK_CHAT="$HOME/.npm-global/bin/kichatty '%n'"
# https://github.com/rodrigo-sys/kichatty

# YOUTUBE_CHAT="pytchatty '%u'"
# https://github.com/rodrigo-sys/pytchatty

# TWITCH_CHAT="chatterino -c '%n'"

# If a platform does not has a custom chat, it will use the browser.
# To specify a browser other than your default, use the TWC_BROWSER option
# placeholders: %u - represents the url
# TWC_BROWSER="qutebrowser --target window %u"

# If %u is missing, the url will be appended to the end
# TWC_BROWSER="qutebrowser --target window"
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
