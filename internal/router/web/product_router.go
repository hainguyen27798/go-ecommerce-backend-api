package web

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/middlewares"
)

type ProductRouter struct{}

func (productRouter *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public router
	productRouterPublic := Router.Group("/products")
	{
		productRouterPublic.GET("")
		productRouterPublic.GET(":id")
	}

	// private router
	productRouterPrivate := Router.Group("/products")
	productRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		productRouterPrivate.POST("")
		productRouterPrivate.PATCH(":id")
	}
}
