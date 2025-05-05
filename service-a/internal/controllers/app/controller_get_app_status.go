package app

import (
	"net/http"

	"github.com/hydde7/goexpert-challenge-2/service-a/internal/handler"
)

type ControllerGetAppStatus struct {
	handler.TransactionControllerImpl
}

type appstatus struct {
	Status string `json:"status"`
}

func (c *ControllerGetAppStatus) Execute(payload interface{}) (result handler.ResponseController) {
	result = handler.NewJsonResponseController()
	result.SetResult(http.StatusOK, appstatus{Status: "OK"})
	return
}

func NewControllerGetAppStatus() handler.Controller {
	return &ControllerGetAppStatus{}
}
