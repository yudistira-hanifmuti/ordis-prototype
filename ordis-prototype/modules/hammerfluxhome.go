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
	fetchEntry.Info("appended some data")
}
