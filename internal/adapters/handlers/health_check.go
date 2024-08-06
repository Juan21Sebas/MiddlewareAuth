package handlers

import (
	"net/http"

	"MiddlewareAuth/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type HealthHttp struct{}

func MakeNewHealthController() HealthHttp {
	return HealthHttp{}
}

func (h *HealthHttp) HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, domain.Health{Status: "UP"})
	}
}
