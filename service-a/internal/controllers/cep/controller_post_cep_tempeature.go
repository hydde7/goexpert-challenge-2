package cep

import (
	"net/http"

	"github.com/hydde7/goexpert-challenge-2/service-a/internal/cfg"
	"github.com/hydde7/goexpert-challenge-2/service-a/internal/handler"
	"github.com/hydde7/goexpert-challenge-2/service-a/internal/models"
	"github.com/hydde7/goexpert-challenge-2/service-a/internal/utils"
	"go.opentelemetry.io/otel"
)

type ControllerGetCepTemperature struct {
	handler.TransactionControllerImpl
}

func (c *ControllerGetCepTemperature) Execute(payload interface{}) (result handler.ResponseController) {
	result = handler.NewJsonResponseController()
	var cep *models.CepPayload

	cep, ok := payload.(*models.CepPayload)
	if !ok {
		result.SetResult(http.StatusUnprocessableEntity, "invalid zipcode")
		return
	}

	if !utils.ValidateCEP(cep.Cep) {
		result.SetResult(http.StatusUnprocessableEntity, "invalid zipcode")
		return
	}

	tracer := otel.Tracer(cfg.Otl.ServiceName)
	ctx, span := tracer.Start(c.GetContext(), "Call-Service-B")
	defer span.End()

	var resp models.CepResponse
	httpResp, err := utils.GetJSON(ctx, cfg.App.ServiceBURL+"/"+cep.Cep, &resp)
	if err != nil {
		result.SetResult(http.StatusInternalServerError, "error calling service B")
		return
	}

	if httpResp.StatusCode != http.StatusOK {
		result.SetResult(http.StatusInternalServerError, "error calling service B")
		return
	}

	result.SetResult(http.StatusOK, resp)
	return
}

func NewControllerGetCepTemperature() handler.Controller {
	return &ControllerGetCepTemperature{}
}
