package cfg

var App = &appConfig{}

type appConfig struct {
	Environment string `json:"APP_ENV"`
	LogLevel    string `json:"APP_LOG_LEVEL"`
}
