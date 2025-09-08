package main

import (
	"analytics-layer/config"
	"analytics-layer/public"
	"os"
	"os/signal"
	"syscall"

	"github.com/HOA-Share/hoashare-go-common/http_server"
	"github.com/HOA-Share/hoashare-go-common/logger"

)

func main(){
	logger.GetLogger().Sugar().Info()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	app := http_server.NewRouter(http_server.RouterConfig{
	AppName:           "uber-data-pipeline",
	StrictRouting:     true,
	IsRecoveryEnabled: true,
	IsLoggerEnabled:   true,
	Port:              config.GetConifg().ServerConfig.Port,
	GlobalCors:        []string{},
	})
	routes.MountRoutes(app)
	
	config.Init()

	app.StartAsync()
	logger.GetLogger().Sugar().Info("server has started")
	<-sig
}