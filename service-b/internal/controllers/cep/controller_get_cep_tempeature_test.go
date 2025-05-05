package cep

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hydde7/goexpert-challenge-2/service-b/internal/cfg"
	"github.com/stretchr/testify/assert"
)

var apiKey = "COLOCAR A API KEY AQUI PARA TESTES"

func TestControllerGetCepTemperature_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	cfg.FreeWeather.ApiKey = apiKey

	r.GET("/cep/:cep", HandleGetCepTemperature)

	req, _ := http.NewRequest("GET", "/cep/19050000", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestControllerGetCepTemperature_InvalidCep(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	cfg.FreeWeather.ApiKey = apiKey

	r.GET("/cep/:cep", HandleGetCepTemperature)

	req, _ := http.NewRequest("GET", "/cep/123456", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, 442, resp.Code)
}

func TestControllerGetCepTemperature_InvalidApiKey(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	cfg.FreeWeather.ApiKey = "invalido!"

	r.GET("/cep/:cep", HandleGetCepTemperature)

	req, _ := http.NewRequest("GET", "/cep/19050000", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, 500, resp.Code)
	cfg.FreeWeather.ApiKey = apiKey
}

func TestControllerGetCepTemperature_CepNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	cfg.FreeWeather.ApiKey = apiKey

	r.GET("/cep/:cep", HandleGetCepTemperature)

	req, _ := http.NewRequest("GET", "/cep/00000000", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, 404, resp.Code)
}
