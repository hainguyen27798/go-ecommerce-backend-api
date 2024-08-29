package global

import (
	"database/sql"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/pkg/logger"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var (
	Config setting.Config
	Logger *logger.Zap
	Mdb    *sql.DB
	Rdb    *redis.Client
)
