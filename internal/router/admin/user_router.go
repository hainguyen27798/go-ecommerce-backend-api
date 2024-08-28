package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/middlewares"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/wires"
)

type UserRouter struct{}

func (userRouter *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController, _ := wires.InitUserRouterHandler()

	// private router
	userRouterPrivate := Router.Group("/users")
	userRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		userRouterPrivate.GET("", userController.GetAllUsers)
		userRouterPrivate.GET(":id")
		userRouterPrivate.DELETE(":id")
	}
}
