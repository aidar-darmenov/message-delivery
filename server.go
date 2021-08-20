package main

import (
	"github.com/aidar-darmenov/message-delivery/config"
	"github.com/aidar-darmenov/message-delivery/service"
	"github.com/aidar-darmenov/message-delivery/webservice"
	"go.uber.org/zap"
	"log"
)

func main() {

	cfg := config.NewConfiguration("/config.config.json")

	var loggerConfig = zap.NewProductionConfig()
	loggerConfig.Level.SetLevel(zap.DebugLevel)

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	// Creating abstract service(business logic) layer
	s := service.NewService(&cfg, logger)

	// Creating abstract webService(delivery) layer
	ws := webservice.NewWebService(s)
	ws.Start()
}
