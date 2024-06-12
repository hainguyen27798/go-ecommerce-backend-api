package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		Host     string `mapstructure:"host"`
		Username string `mapstructure:"user"`
		Password string `mapstructure:"pass"`
		Database string `mapstructure:"db_name"`
	} `mapstructure:"databases"`
}

func main() {
	v := viper.New()
	v.AddConfigPath("./configs/")
	v.SetConfigName("local")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}
	fmt.Println(config)
}
