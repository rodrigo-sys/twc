# twc

<img src="https://github.com/user-attachments/assets/616a846f-b76d-496f-bf36-096c1bb1e954" height=300/>
<details>
  <summary>vods</summary>
  <img src="https://github.com/user-attachments/assets/6e9362e3-2b95-4a69-836f-978dfaf664db" height=300/> 
</details>

## Overview
  
🎥 **TWC** is a Terminal User Interface (TUI) application that allows you to watch Video on Demand
 (VOD) and live streams from your favorite streamers on Kick, YouTube, and Twitch. 🚀 Upon
 launching, you'll be greeted with a menu displaying a list of channels 📺, enabling you to view
 previous streams ⏪ or watch their current live broadcasts 📡 in a local video player 🎬, along
 with a custom chat client 💬 or the official pop-out chat 🗨.

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
npm install -g 'https://github.com/rodrigo-sys/stealth-cli'
```
```sh
sudo pacman -S npm --noconfirm
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
go install 'https://github.com/rodrigo-sys/twc'
```

### Usage
Just run 
``` sh
twc
```  
to open the interface.

at you first run it will ask you to add some _channels_  
![2025-01-25_12:45](https://github.com/user-attachments/assets/d6e3a650-d12a-458a-9e2c-cc320bfa462b)

follow the screen instructions to open the channels file  
here is an example channels file:
```
#syntax:
#<username> <spaces or tab> <platform>

xQc	twitch
Markiplier	youtube
Trainwreckstv	kick
Ninja	twitch
Jacksepticeye	youtube
Sodapoppin	kick
Asmongold	twitch
CorpseHusband	youtube
```

**flags**
```
Usage of twc:
  -e	open channels file in default text editor
  -o string
    	open channel
  -v string
    	view vods of channel
```


**🚧 Work in Progress 🚧**
 
 This README is still being developed. Please check back later for more information.




