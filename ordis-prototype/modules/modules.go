package modules

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"app/ordis-prototype/config"

	"github.com/robfig/cron"
)

type ModuleManager struct {
	BaseConfig *config.BaseConfiguration
	Cron       *cron.Cron
	Logger     *logrus.Logger
}

func NewModuleManager(log *logrus.Logger, conf *config.BaseConfiguration) *ModuleManager {
	log.WithFields(logrus.Fields{
		"Context": "ordis/modules/modules.go",
	}).Info("ModuleManager is initialized.")
	moduleCron := cron.New()
	manager := &ModuleManager{conf, moduleCron, log}
	return manager
}

func (thisObject *ModuleManager) PrepareModule() {

	logEntry := thisObject.Logger.WithFields(logrus.Fields{
		"Context": "ordis/modules/modules.go",
	})
	logEntry.Info("Preparing modules.")

	//retrieve installed modules
	modules := thisObject.BaseConfig.InstalledModules

	//Add fetching to cron
	for _, module := range modules {
		thisObject.Logger.WithFields(logrus.Fields{
			"Context":    "ordis/modules/modules.go",
			"ModuleName": module,
		}).Info("Adding fetch schedule.")
		var jobFunc = FetchHandlers[module]
		thisObject.Cron.AddFunc("@every 1m", func() {
			jobFunc.(func())()
		})
	}
	successMsg := fmt.Sprintf("All %d modules are ready.", cap(modules))
	fmt.Println(successMsg)
	fmt.Println("Data fetching has started. Please see the log file.")
	logEntry.Info(successMsg)

	go thisObject.Cron.Start()
}

func (thisObject *ModuleManager) getModuleConfiguration(moduleName string) *config.ModuleConfiguration {
	var moduleConfig *config.ModuleConfiguration
	//retrieve config directory
	configDir := thisObject.BaseConfig.ConfigPath + "modules/"

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
