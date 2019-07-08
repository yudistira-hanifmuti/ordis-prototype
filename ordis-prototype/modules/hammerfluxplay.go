package modules

import "github.com/sirupsen/logrus"

func init() {

	//register module to module handler
	registerFetchHandler("hammerfluxplay", HammerfluxPlayFetch)
}

func HammerfluxPlayFetch() {
	fetchEntry := ObjModuleManager.Logger.WithFields(logrus.Fields{
		"Context": "ordis/modules/hammerfluxplay.go",
		"Event":   "Fetching Handling",
	})
	fetchEntry.Info("appended some data")
}
