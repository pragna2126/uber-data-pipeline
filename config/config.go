package config

import (
"github.com/HOA-Share/hoashare-go-common/config"
)

type Settings struct {
	DBConfig config.DatabaseConfig `prefix:"DB"`
	ServerConfig config.ServerConfig
}

var settings Settings

func Init() {
	err := config.LoadConfig(&settings)
	if err != nil {
		panic(err)
	}
	ConnectDB()
	
}

func GetConifg() *Settings {
	return &settings
}