package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/middlewares"
)

type AuthenticateRouter struct{}

func (authenticateRouter *AuthenticateRouter) InitAuthenticateRouter(Router *gin.RouterGroup) {
	// public router
	authenticateRouterPublic := Router.Group("/auth")
	{
		authenticateRouterPublic.POST("/register")
		authenticateRouterPublic.POST("/login")
	}

	// private router
	authenticateRouterPrivate := Router.Group("")
	authenticateRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		authenticateRouterPrivate.GET("/me")
	}
}
