package initialize

import (
	"github.com/gin-gonic/gin"
	c "github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/controllers"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/middlewares"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthMiddleware())
	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			user.GET("", c.NewUserController().GetAllUsers)
		}
	}
	return r
}
