package initialize

import (
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/global"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
