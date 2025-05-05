package cfg

var FreeWeather = &freeWeatherConfig{}

type freeWeatherConfig struct {
	ApiKey string `json:"FREEWEATHER_API_KEY"`
}
