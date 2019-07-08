package modules

import (
	"encoding/json"
	"fmt"
	"os"

	"app/ordis-prototype/config"

	"github.com/robfig/cron"
)

type ModuleManager struct {
	BaseConfig *config.BaseConfiguration
	Cron       *cron.Cron
}

func NewModuleManager(conf *config.BaseConfiguration) *ModuleManager {
	fmt.Println("\tCreating new cron object.")
	moduleCron := cron.New()
	manager := &ModuleManager{conf, moduleCron}
	return manager
}

func (thisObject *ModuleManager) PrepareModule() {
	//retrieve installed modules
	modules := thisObject.BaseConfig.InstalledModules

	//Add fetching to cron
	for _, module := range modules {
		fmt.Printf("\t\t~~~ Start adding %s fetch schedule\n", module)
		var jobFunc = FetchHandlers[module]
		thisObject.Cron.AddFunc("@every 1m", func() {
			jobFunc.(func())()
		})
	}

	go thisObject.Cron.Start()
}

func (thisObject *ModuleManager) getModuleConfiguration(moduleName string) *config.ModuleConfiguration {
	var moduleConfig *config.ModuleConfiguration
	//retrieve config directory
	configDir := thisObject.BaseConfig.ConfigDir + "modules/"

	//read base configuration file
	file, err := os.Open(configDir + moduleName + ".json")
	if err != nil {
		fmt.Println("Error when opening JSON config file: ", err)
		os.Exit(-2)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&moduleConfig)
	if err != nil {
		fmt.Println("Error when decoding JSON config file: ", err)
		os.Exit(-3)
	}

	return moduleConfig
}

func ScanModules(modules []string) {
	fmt.Println("~~~ Start assigning schedule ~~~")
	fmt.Println()
	for _, module := range modules {
		fmt.Printf("~~~ Start adding %s fetch schedule\n", module)
		//do something
		fmt.Printf("~~~ Start adding %s predict schedule\n", module)
		//do something

	}
}
