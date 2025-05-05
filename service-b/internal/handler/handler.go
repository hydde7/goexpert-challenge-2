package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestWithController(c *gin.Context, payload interface{}, controller Controller) {
	if payload != nil {
		err := c.ShouldBindJSON(payload)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	controller.SetDefaultLog(logrus.WithFields(logrus.Fields{
		"action": c.Request.Method + " " + c.Request.URL.Path,
	}))
	controller.SetContext(c)
	controller.SetParams(c.Params)

	result := controller.Execute(payload)
	status := result.GetStatusCode()

	if result.IsAbort() {
		if err := result.GetErrors(); len(err) > 0 {
			c.AbortWithStatusJSON(status, gin.H{"error": err})
		} else {
			c.AbortWithStatus(status)
		}
		return
	}

	result.Write(c)
}
