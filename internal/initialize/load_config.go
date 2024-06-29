package initialize

import (
	"fmt"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	v := viper.New()
	v.AddConfigPath("./configs/")
	v.SetConfigName("local")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
	fmt.Println(global.Config)
}
