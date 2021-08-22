package service

import (
	"github.com/aidar-darmenov/message-delivery/config"
	"github.com/aidar-darmenov/message-delivery/interfaces"
	"github.com/aidar-darmenov/message-delivery/model"
	"go.uber.org/zap"
	"sync"
)

type Service struct {
	Configuration   interfaces.Configuration
	Clients         model.Clients // Using sync.Map to store connected clients
	Logger          *zap.Logger
	ChannelMessages chan model.MessageToClients
}

func NewService(cfg *config.Configuration, logger *zap.Logger) interfaces.Service {
	//Here can be any other objects like DB, Cache, any kind of delivery services

	channelMessages := make(chan model.MessageToClients, cfg.ChannelMessagesSize)

	return &Service{
		Configuration: cfg,
		Logger:        logger,
		Clients: model.Clients{
			Map: &sync.Map{},
		},
		ChannelMessages: channelMessages,
	}
}

func (s *Service) GetLogger() *zap.Logger {
	return s.Logger
}

func (s *Service) GetConfigParams() *config.Configuration {
	return s.Configuration.Params()
}

func (s *Service) GetClients() model.Clients {
	return s.Clients
}
