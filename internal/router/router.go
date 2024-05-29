package router

import (
	"github.com/gin-gonic/gin"
	c "github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/controllers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			user.GET("", c.NewUserController().GetAllUsers)
		}
	}
	return r
}
