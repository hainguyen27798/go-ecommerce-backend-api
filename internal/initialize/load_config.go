package initialize

import (
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/global"
	"github.com/spf13/viper"
	"os"
)

func LoadConfig() {
	name := "local"

	if os.Getenv("MODE") == "prod" {
		name = "production"
	}

	v := viper.New()
	v.AddConfigPath("./configs/")
	v.SetConfigName(name)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
}
