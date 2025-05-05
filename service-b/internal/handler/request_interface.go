package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Controller interface {
	SetParams(params gin.Params)
	SetContext(context *gin.Context)
	SetDefaultLog(log *logrus.Entry)
	GetParam(key string) string
	GetQueryParam(key string) string
	GetContext() *gin.Context
	Execute(payload interface{}) ResponseController
}

type TransactionControllerImpl struct {
	Params     gin.Params
	context    *gin.Context
	DefaultLog *logrus.Entry
}

func (t *TransactionControllerImpl) SetParams(params gin.Params) {
	t.Params = params
}

func (t *TransactionControllerImpl) SetContext(context *gin.Context) {
	t.context = context
}

func (t *TransactionControllerImpl) SetDefaultLog(log *logrus.Entry) {
	t.DefaultLog = log
}

func (t *TransactionControllerImpl) GetParam(key string) string {
	return t.Params.ByName(key)
}

func (t *TransactionControllerImpl) GetQueryParam(key string) string {
	return t.context.Query(key)
}

func (t *TransactionControllerImpl) GetContext() *gin.Context {
	return t.context
}
