# twc

<img src="https://github.com/user-attachments/assets/616a846f-b76d-496f-bf36-096c1bb1e954" height=300/>
<details>
  <summary>vods</summary>
  <img src="https://github.com/user-attachments/assets/6e9362e3-2b95-4a69-836f-978dfaf664db" height=300/> 
</details>

## Overview  
**TWC** is a TUI application that allows you to watch 🍿 VODs and live streams
from your favorite streamers of 
Kick <img src="https://github.com/user-attachments/assets/d0e305b3-5a25-432f-b9b2-9f0b8a965de5" alt="Description" width="15"/>,
YouTube <img src="https://github.com/user-attachments/assets/d104ff17-46c3-4817-9875-a4e13d81f00c" alt="Description" width="17"/>,
and Twitch <img src="https://github.com/user-attachments/assets/c3d44e61-1a60-4b4a-bf87-14f404d65390" alt="Description" width="15"/>.
Upon launching 🚀 , you'll be greeted with a menu displaying a list of channels 📺, enabling you to view
 previous streams ⏪ or watch their current live broadcasts 📡
in a local video player 🎬, along with the official pop-out chat 🗨 or a custom chat client 💬.

 



### Installation
**requirements**
<details>
  <summary><a href="https://mpv.io/">mpv</a></summary>
  
```sh
sudo apt install mpv -y
```
```sh
sudo pacman -S mpv --noconfirm
```
</details>
<details>
  <summary><a href="https://github.com/yt-dlp/yt-dlp">yt-dlp</a></summary>
  
```sh
sudo apt install yt-dlp -y
```
```sh
sudo pacman -S yt-dlp --noconfirm
```
</details>
<details>
  <summary><a href="https://github.com/rodrigo-sys/stealth-cli">stealth-cli</a> (for kick support)</summary>
  
```sh
sudo apt install npm -y
sudo apt install libnss3 -y
npm config set prefix '~/.local'
npm install -g 'https://github.com/rodrigo-sys/stealth-cli'
```
```sh
sudo pacman -S npm --needed --noconfirm
npm config set prefix '~/.local'
npm install -g 'https://github.com/rodrigo-sys/stealth-cli'
```
</details>
<details>
  <summary><a href="https://go.dev/">go</a></summary>
  
```sh
sudo apt install golang -y
```
```sh
sudo pacman -S go --noconfirm
```
</details>

**the program**
```sh
git clone 'https://github.com/rodrigo-sys/twc' /tmp/twc
(cd /tmp/twc ; go build -o ~/.local/bin/twc main.go)
```

### Usage
``` sh
twc # to open the user interface
```  
**navigation**  
`j` next item  
`k` previous item  
`l` select current item  
`h` go back to main menu  

**current behaviour**  
when you select an item
- If the channel is live, it will open the stream in mpv and the chat in the official pop-out. 
- If the channel is offline, it will open a menu displaying their VODs.

**flags**
```
Usage of twc:
  -e	open channels file in default text editor
  -o string
    	open channel
  -v string
    	view vods of channel
```
examples:
```
twc -o 'Markiplier youtube' # open live stream
twc -v 'Ninja twitch' # open VODs menu

# you can omit the platform if the channels is in your channels file
twc -o xQc 
twc -v Asmongold
```

**🚧 Work in Progress 🚧**
 
 This README is still being developed. Please check back later for more information.




