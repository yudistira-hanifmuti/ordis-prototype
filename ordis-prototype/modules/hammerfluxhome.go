package modules

import (
	"github.com/sirupsen/logrus"
)

func init() {

	//register module to module handler
	registerFetchHandler("hammerfluxhome", HammerfluxHomeFetch)
}

func HammerfluxHomeFetch() {
	fetchEntry := ObjModuleManager.Logger.WithFields(logrus.Fields{
		"Context": "ordis/modules/hammerfluxhome.go",
		"Event":   "Fetching Handling",
	})

	moduleConfig := ObjModuleManager.getModuleConfiguration("hammerfluxhome")
	features := moduleConfig.Features

	fetchEntry.Info("Try fetch for main feature ", features[0], ".")

	for _, feature := range features[1:] {
		fetchEntry.Info("Fetch ", feature, " feature.")
	}

	fetchEntry.Info("appended some data")
}
