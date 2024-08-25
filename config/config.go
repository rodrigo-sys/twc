package config

import (
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func getConfigPath() string {
	var default_config_path string
	var config_path string

	if os.Getenv("TWC_CONFIG_PATH") == "" {
		home, _ := os.UserHomeDir()
		default_config_path = home + "/.config/twc/config.env"

		config_path = default_config_path
	} else {
		config_path = os.Getenv("TWC_CONFIG_PATH")
	}

	/* td: check if config_path is a fullpath
	or think what to do if it is relative */

	return config_path
}

func init() {
	config_path := getConfigPath()
	godotenv.Load(config_path)
}
