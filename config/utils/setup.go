package config

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func createConfig() {
	var source_file io.Reader
	example_config_path := os.Getenv("TWC_EXAMPLE_CONFIG_PATH")

	if example_config_path == "" {
		url := "https://raw.githubusercontent.com/rodrigo-sys/telescope-mantoc/main/.env.example"
		response, _ := http.Get(url)
		defer response.Body.Close()
		source_file = response.Body
		// defer source_file.Close()
		// td: replace <user> wih os.UserConfigDir
	} else {
		source_file, _ = os.Open(example_config_path)
		// defer source_file.Close()
	}

	destination_file, _ := os.Create(GetConfigPath())
	defer destination_file.Close()
	io.Copy(destination_file, source_file)
}

func SetupConfig() {
	/*td: think to do this in a setup or installer script just one time*/

	// create config folder
	config_path := GetConfigPath()
	os.MkdirAll(filepath.Dir(config_path), os.ModePerm)

	// create config file
	if _, err := os.Stat(config_path); os.IsNotExist(err) {
		createConfig()
	}
}
