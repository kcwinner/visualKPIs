package common

import (
	"flag"
)

var configFile string

// StartUp Initialze the server and app config
func StartUp() {
	loadParams()
	initConfig(configFile)
	setupLogging()
}

func loadParams() {
	flag.StringVar(&configFile, "configFile", "common/config.json", "location of the config")
	flag.Parse()
}
