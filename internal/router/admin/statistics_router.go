package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/middlewares"
)

type StatisticsRouter struct{}

func (statisticsRouter *StatisticsRouter) InitStatisticsRouter(Router *gin.RouterGroup) {
	// private router
	statisticsRouterPrivate := Router.Group("/statistics")
	statisticsRouterPrivate.Use(middlewares.AuthMiddleware())
	{
		statisticsRouterPrivate.GET("/users")
		statisticsRouterPrivate.GET("/products")
	}
}
