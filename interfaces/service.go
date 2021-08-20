package interfaces

import (
	"github.com/aidar-darmenov/message-delivery/config"
	"go.uber.org/zap"
	"sync"
)

type Service interface {
	GetLogger() *zap.Logger
	GetConfigParams() *config.Configuration
	GetConnectedClients() sync.Map
}
