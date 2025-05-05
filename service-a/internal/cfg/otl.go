package cfg

var Otl = &OtlConfig{}

type OtlConfig struct {
	OTELPEndpoint string `json:"OTEL_ENDPOINT"`
	ServiceName   string `json:"OTEL_SERVICE_NAME"`
}
