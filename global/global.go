package global

import (
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/pkg/logger"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.Zap
)
