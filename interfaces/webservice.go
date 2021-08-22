package interfaces

import "github.com/gin-gonic/gin"

type WebService interface {
	Start()
	GetConnectedClientsIds(*gin.Context)
	SendMessageToClientsByIds(*gin.Context)
}
