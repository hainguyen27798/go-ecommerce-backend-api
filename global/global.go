package global

import (
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/pkg/logger"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.Zap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
