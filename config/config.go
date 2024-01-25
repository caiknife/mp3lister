package config

var (
	Config *AppConfig
)

func init() {
	initConfig()
}

func initConfig() {
	Config = &AppConfig{}
	initYAMLConfig(Config, "config.yml")
}
