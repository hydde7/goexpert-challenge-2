package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hydde7/goexpert-challenge-2/service-b/internal/handler"
)

// GET Endpoints

// [GET] /appstatus
// @Summary Get app status
// @Tags App
// @Description Get app status
// @Accept json
// @Produce json
// @Success 200 {object} string "App Status"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /appstatus [get]
func HandleGetAppStatus(c *gin.Context) {
	handler.RequestWithController(c, nil, NewControllerGetAppStatus())
}
