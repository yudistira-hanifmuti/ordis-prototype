package config

type BaseConfiguration struct {
	InstalledModules []string `json:"InstalledModules"`
	ConfigPath       string   `json:"ConfigPath"`
	LogPath          string   `json:"LogPath"`
}
