package Config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Database struct {
		UserName string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"database"`
	Port string `json:"port"`
}

func ParseConfigFile() Configuration {
	os.Chdir("Config")
	var config Configuration
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		panic(err)
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
