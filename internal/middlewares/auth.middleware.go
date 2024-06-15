package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/pkg/response"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if !strings.HasPrefix(token, "Bearer ") {
			response.ErrorResponse(c, response.ErrInvalidToken)
			c.Abort()
			return
		}
		c.Next()
	}
}
