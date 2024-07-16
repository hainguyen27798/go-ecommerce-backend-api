package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/middlewares"
)

type UserRouter struct{}

func (userRouter *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// private router
	userRouterPrivate := Router.Group("/users")
	userRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		userRouterPrivate.GET("")
		userRouterPrivate.GET(":id")
		userRouterPrivate.DELETE(":id")
	}
}
