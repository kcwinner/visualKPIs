package common

import (
	"fmt"
	"log"
	"os"
)

func setupLogging() {
	logFile := AppConfig.LogFile
	if logFile == "" {
		return
	}

	if !FileExists(logFile) {
		nf, err := os.Create(logFile)
		if err != nil {
			fmt.Println("Error - Cannot create log file at " + logFile)
		}
		nf.Close()
	}

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		fmt.Println("Error - Cannot get to log file at " + logFile)
		panic(err)
	}

	log.SetOutput(f)
	log.Print("---")
}
