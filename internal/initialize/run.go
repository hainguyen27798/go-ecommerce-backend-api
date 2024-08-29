package initialize

import (
	"fmt"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/global"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	err := r.Run(fmt.Sprintf(":%d", global.Config.Server.Port))
	if err != nil {
		return
	}
}
