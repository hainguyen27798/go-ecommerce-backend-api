package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
