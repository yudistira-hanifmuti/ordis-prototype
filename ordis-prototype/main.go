package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"app/ordis-prototype/config"
	"app/ordis-prototype/modules"

	log "github.com/sirupsen/logrus"
)

//configuration
var configFile = "base"
var baseConfig config.BaseConfiguration

const configDir = "/home/nakama/Documents/go/src/app/ordis/"

//logger
var logger = log.New()

func init() {

	//read base configuration file
	file, err := os.Open(configDir + configFile + ".json")
	if err != nil {
		fmt.Println("Error when opening JSON config file: ", err)
		os.Exit(-2)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&baseConfig)
	if err != nil {
		fmt.Println("Error when decoding JSON config file: ", err)
		os.Exit(-3)
	}

	//setup logger
	logFile := baseConfig.LogPath + "ordis-module.log"
	file, err = os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("File logger not found. Creating log file.")
		file, err = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY, 0666)
	}
	logger.SetOutput(file)
	logger.Info("File logger is initialized.")

}

func main() {
	//initialize installed module
	modules.ObjModuleManager = modules.NewModuleManager(logger, &baseConfig)
	modules.ObjModuleManager.PrepareModule()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
