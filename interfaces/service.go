package interfaces

import (
	"github.com/aidar-darmenov/message-delivery/config"
	"github.com/aidar-darmenov/message-delivery/model"
	"go.uber.org/zap"
)

type Service interface {
	GetLogger() *zap.Logger
	GetClients() model.Clients
	GetConfigParams() *config.Configuration
	SendMessageToClientsByIds(message model.MessageToClients) *model.Exception
}
