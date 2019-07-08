package modules

import "github.com/sirupsen/logrus"

func init() {

	//register module to module handler
	registerFetchHandler("regressiontest", RegressionTestFetch)
}

func RegressionTestFetch() {
	fetchEntry := ObjModuleManager.Logger.WithFields(logrus.Fields{
		"Context": "ordis/modules/regressiontest.go",
		"Event":   "Fetching Handling",
	})
	fetchEntry.Info("appended some data")
}
