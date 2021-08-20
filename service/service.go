package service

import (
	"github.com/aidar-darmenov/message-delivery/config"
	"github.com/aidar-darmenov/message-delivery/interfaces"
	"github.com/aidar-darmenov/message-delivery/model"
	"go.uber.org/zap"
	"sync"
)

type Service struct {
	Configuration interfaces.Configuration
	Clients       model.Clients // Using sync.Map to store connected clients
	Logger        *zap.Logger
}

func NewService(cfg *config.Configuration, logger *zap.Logger) *Service {
	//Here can be any other objects like DB, Cache, any kind of delivery services
	return &Service{
		Configuration: cfg,
		Logger:        logger,
		Clients: model.Clients{
			Map: &sync.Map{},
			Ids: nil,
		},
	}
}

func (s *Service) GetLogger() *zap.Logger {
	return s.Logger
}

func (s *Service) GetConfigParams() *config.Configuration {
	return s.Configuration.Params()
}
