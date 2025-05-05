package cfg

import (
	"github.com/urfave/cli"
)

var Flags = []cli.Flag{
	// App flags
	&cli.StringFlag{
		Name:        "app-env",
		Destination: &App.Environment,
		EnvVar:      "APP_ENV",
		Value:       Development,
	},
	&cli.StringFlag{
		Name:        "app-log-level",
		Destination: &App.LogLevel,
		EnvVar:      "APP_LOG_LEVEL",
		Value:       "debug",
	},
	&cli.StringFlag{
		Name:        "service-b-url",
		Destination: &App.ServiceBURL,
		EnvVar:      "SERVICE_B_URL",
		Value:       "http://localhost:9090",
	},
	// OpenTelemetry flags
	&cli.StringFlag{
		Name:        "otl-endpoint",
		Destination: &Otl.OTELPEndpoint,
		EnvVar:      "OTEL_ENDPOINT",
		Value:       "localhost:4317",
	},
	&cli.StringFlag{
		Name:        "otl-service-name",
		Destination: &Otl.ServiceName,
		EnvVar:      "OTEL_SERVICE_NAME",
		Value:       "service-b",
	},
}
