package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"app/ordis-prototype/config"
	"app/ordis-prototype/modules"
)

//configuration
var configFile = "base"
var baseConfig config.BaseConfiguration

const configDir = "/home/nakama/Documents/go/src/app/ordis/"

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

}

func main() {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("Preparing internal...")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println()

	fmt.Println("\tCreating Module Manager object.")
	manager := modules.NewModuleManager(&baseConfig)
	modules.BaseConfiguration = &baseConfig
	fmt.Println("\t~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("\tPreparing modules...")
	manager.PrepareModule()
	fmt.Println("\t~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	fmt.Println("\tAll modules are ready.")
	fmt.Println("\tStart fetching data. See log.")
	fmt.Println("\t~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
