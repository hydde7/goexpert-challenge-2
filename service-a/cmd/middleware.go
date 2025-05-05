package cmd

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Middleware struct{}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) Options(c *gin.Context) {
	origin := c.GetHeader("Origin")

	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Max-Age", "3600")
	c.Header("Access-Control-Allow-Methods", "GET, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Encoding, Authorization")
	c.Header("Access-Control-Allow-Credentials", "true")

	c.Header("Content-Type", "application/json")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}

func (a *Middleware) Errors(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(-1, c.Errors)
	}
}

func (a *Middleware) NotFound(c *gin.Context) {
	c.Error(errors.New("not found"))
	c.AbortWithStatus(http.StatusNotFound)
}

func (a *Middleware) MethodNotAllowed(c *gin.Context) {
	c.Error(errors.New("method not allowed"))
	c.AbortWithStatus(http.StatusMethodNotAllowed)
}
