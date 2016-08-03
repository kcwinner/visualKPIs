package common

import (
	"encoding/json"
	"os"
)

type configuration struct {
	DatabaseServer string `json:"databaseserver"`
	Database       string `json:"database"`
	LogFile        string `json:"logfile"`
}

// AppConfig Struct used to access the configuration data
var AppConfig configuration

func initConfig(configFile string) {
	loadAppConfig(configFile)
}

func loadAppConfig(configFile string) {
	file, err := os.Open(configFile)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		panic(err)
	}
}

// FileExists Checks if the file exists
func FileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
