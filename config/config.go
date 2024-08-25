package config

import (
	config "twc/config/utils"

	"github.com/joho/godotenv"
)

func init() {
	config_path := config.GetConfigPath()
	godotenv.Load(config_path)
}
