package interfaces

import "github.com/aidar-darmenov/message-delivery/config"

type Configuration interface {
	InitConfigParams(string)
	Params() *config.Configuration
}
