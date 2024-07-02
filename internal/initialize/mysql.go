package initialize

import (
	"fmt"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/global"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/models"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func CheckErrorPanic(err error, errMsg string) {
	if err != nil {
		global.Logger.Error(errMsg, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	config := global.Config.Mysql

	// sql config connection
	dsnFormat := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		dsnFormat,
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	CheckErrorPanic(err, "Init mysql failed")
	global.Logger.Info("Initializing mysql successfully")
	global.Mdb = db

	setPool()
	migrateTables()
}

func setPool() {
	config := global.Config.Mysql
	sqlDB, err := global.Mdb.DB()
	CheckErrorPanic(err, "Mysql error")

	sqlDB.SetConnMaxIdleTime(time.Duration(config.MaxIdleConn))
	sqlDB.SetMaxOpenConns(config.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifeTime))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&models.User{},
		&models.Role{},
	)
	if err != nil {
		global.Logger.Error(err.Error())
	}
}
