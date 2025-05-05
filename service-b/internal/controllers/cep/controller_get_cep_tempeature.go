package cep

import (
	"net/http"

	"github.com/hydde7/goexpert-challenge-2/service-b/internal/cfg"
	"github.com/hydde7/goexpert-challenge-2/service-b/internal/handler"
	"github.com/hydde7/goexpert-challenge-2/service-b/internal/services"
	"github.com/hydde7/goexpert-challenge-2/service-b/internal/utils"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

type ControllerGetCepTemperature struct {
	handler.TransactionControllerImpl
}

func (c *ControllerGetCepTemperature) Execute(payload interface{}) (result handler.ResponseController) {
	result = handler.NewJsonResponseController()
	cep := c.GetParam("cep")
	logrus.Infof("cep: %s", cep)
	if !utils.ValidateCEP(cep) {
		logrus.Errorf("invalid zipcode: %s", cep)
		result.SetResult(http.StatusUnprocessableEntity, "invalid zipcode")
		return
	}

	tracer := otel.Tracer(cfg.Otl.ServiceName)
	ctx, span1 := tracer.Start(c.GetContext(), "Lookup-CEP")
	cepPayload, err := services.ViaCepRequest(ctx, cep)
	span1.End()
	if cepPayload.Cep == "" || err != nil {
		logrus.Errorf("error looking up zipcode: %s", cep)
		result.SetResult(http.StatusNotFound, "zipcode not found")
		return
	}

	ctx2, span2 := tracer.Start(ctx, "Fetch-Temperature")
	weatherPayload, err := services.FreeWeatherRequest(ctx2, cepPayload.Localidade)
	span2.End()
	if err != nil {
		logrus.Errorf("error fetching temperature: %s", err)
		result.SetResult(http.StatusInternalServerError, err.Error())
		return
	}

	response := getCepTemperatureResponse{
		City:                  weatherPayload.Location.Name,
		TemperatureCelsius:    weatherPayload.Current.TempC,
		TemperatureFahrenheit: weatherPayload.Current.TempF,
		TemperatureKelvin:     weatherPayload.Current.TempC + 273,
	}

	result.SetResult(http.StatusOK, response)
	return
}

type getCepTemperatureResponse struct {
	City                  string  `json:"city"`
	TemperatureCelsius    float64 `json:"temp_C"`
	TemperatureFahrenheit float64 `json:"temp_F"`
	TemperatureKelvin     float64 `json:"temp_K"`
}

func NewControllerGetCepTemperature() handler.Controller {
	return &ControllerGetCepTemperature{}
}
