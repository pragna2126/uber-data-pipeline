package config

import (
	"github.com/HOA-Share/hoashare-go-common/config"
	"github.com/HOA-Share/hoashare-go-common/db"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"fmt"
)

var DB *gorm.DB

func ConnectDB() {
	dbConn, err := db.ConnectDB(config.DatabaseConfig{
		Host: settings.DBConfig.Host,
		Port: settings.DBConfig.Port,
		Password: settings.DBConfig.Password,
		User: settings.DBConfig.User,
		DBName: settings.DBConfig.DBName,
		SSLMode: settings.DBConfig.SSLMode,
	},(logger.Info)) 
	if err != nil {
		panic(err)
	}
	fmt.Println("Port from config:", settings.DBConfig.Port)

	DB = dbConn.Raw()
}
