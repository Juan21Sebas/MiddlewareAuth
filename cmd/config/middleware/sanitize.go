package middlewares

import (
	URL "net/url"

	"MiddlewareAuth/cmd/config/errors"

	"github.com/gin-gonic/gin"

	"MiddlewareAuth/cmd/entity"
)

func SanitizeRequest() gin.HandlerFunc {
	return func(c *gin.Context) {

		if _, err := URL.Parse(c.Request.URL.Path); err != nil {
			details := make([]entity.Detail, 1)
			details[0].Message = err.Error()
			_ = c.Error(errors.BadRequest(details, errors.CONSUMER))
			c.Abort()
			return
		}
		c.Next()
	}
}
