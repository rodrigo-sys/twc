package setup

import (
	"github.com/joho/godotenv"
)

// func init() {
func LoadConfig() {
	config_path := program_files["config"].GetPath()
	godotenv.Load(config_path)
}
