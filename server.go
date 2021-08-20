package main

import (
	"github.com/aidar-darmenov/message-delivery/config"
	"github.com/aidar-darmenov/message-delivery/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"strconv"
)

func main() {

	cfg := config.NewConfiguration("/config.config.json")

	var loggerConfig = zap.NewProductionConfig()
	loggerConfig.Level.SetLevel(zap.DebugLevel)

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	s := service.NewService(&cfg, logger)

	g := gin.Default()

	g.GET("/clients/connected/all", s.GetAllConnectedClients)

	g.Run(":" + strconv.Itoa(cfg.Params().HttpPort))

}
