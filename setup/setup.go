package setup

import (
	"os"
	"path/filepath"
	filescontent "twc/setup/files-content"
	. "twc/setup/program-file"
)

var program_files = map[string]ProgramFile{
	"config": {
		Path:              os.Getenv("TWC_CONFIG_PATH"),
		Example_file_path: os.Getenv("TWC_EXAMPLE_CONFIG_PATH"),
		Content:           filescontent.Config,
		Name:              "config.env",
		Placeholders: map[string]func() (string, error){
			"<your config dir>":    os.UserConfigDir,
			"<your user home dir>": os.UserHomeDir,
		},
	},
	"channels": {
		Path:              os.Getenv("TWC_CHANNELS_PATH"),
		Example_file_path: os.Getenv("TWC_EXAMPLE_CHANNELS_PATH"),
		Content:           filescontent.Channels,
		Name:              "channels",
	},
}

func CreateProgramFiles() {
	for _, programfile := range program_files {
		path := programfile.GetPath()
		// fmt.Println(path)

		os.MkdirAll(filepath.Dir(path), os.ModePerm)

		if _, err := os.Stat(path); os.IsNotExist(err) {
			programfile.CreateFile()
		}
	}
}
