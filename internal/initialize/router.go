package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/global"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/router"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// Declare router
	adminRouter := router.AppRouter.Admin
	authRouter := router.AppRouter.Auth
	webRouter := router.AppRouter.Web

	mainRouter := r.Group("/v1")
	{
		mainRouter.GET("/check-status")
	}
	{
		adminRouter.InitStatisticsRouter(mainRouter)
		adminRouter.InitUserRouter(mainRouter)
	}
	{
		authRouter.InitAuthenticateRouter(mainRouter)
	}
	{
		webRouter.InitProductRouter(mainRouter)
	}

	return r
}
