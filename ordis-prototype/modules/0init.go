package modules

import (
	"app/ordis-prototype/config"
)

//
var BaseConfiguration *config.BaseConfiguration

//

//Fetch variable to store module handlers
var FetchHandlers map[string]interface{}

func init() {
	FetchHandlers = make(map[string]interface{})
}

func registerFetchHandler(moduleName string, fetchHandlers interface{}) {
	FetchHandlers[moduleName] = fetchHandlers
}
