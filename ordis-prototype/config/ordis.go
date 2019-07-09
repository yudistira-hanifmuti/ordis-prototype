package config

type BaseConfiguration struct {
	InstalledModules []string `json:"InstalledModules"`
	ConfigPath       string   `json:"ConfigPath"`
	LogPath          string   `json:"LogPath"`
}

type ModuleConfiguration struct {
	FeatureFolderLocation string   `json:"FeatureFolderLocation"`
	Features              []string `json:"Features"`
	PredictionExecutable  string   `json:"PredictionExecutable"`
	PredictionFilename    string   `json:"PredictionFilename"`
}
