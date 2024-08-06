package handlers

import (
	"context"
	"net/http"

	"MiddlewareAuth/cmd/utils"

	repository "MiddlewareAuth/internal/core/domain/repository"

	"github.com/gin-gonic/gin"
)

func (h *AuthHttp) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request repository.User
		if err := utils.ShouldBindJSON(c, &request); err != nil {
			return
		}
		response, err := h.service.Login(context.Background(), &request)
		if err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
