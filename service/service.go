package service

import (
	"github.com/aidar-darmenov/message-delivery/config"
	"github.com/aidar-darmenov/message-delivery/interfaces"
	"go.uber.org/zap"
)

type Service struct {
	Configuration interfaces.Configuration
	Logger        *zap.Logger
}

func NewService(cfg *interfaces.Configuration, logger *zap.Logger) *Service {
	//Here can be any other objects like DB, Cache, any kind of delivery services
	return &Service{
		Configuration: *cfg,
		Logger:        logger,
	}
}

func (s *Service) GetLogger() *zap.Logger {
	return s.Logger
}

func (s *Service) GetConfigParams() *config.Configuration {
	return s.Configuration.Params()
}
