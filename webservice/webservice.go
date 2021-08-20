package webservice

import (
	"github.com/aidar-darmenov/message-delivery/interfaces"
	gin_webservice "github.com/aidar-darmenov/message-delivery/webservice/gin"
	"github.com/gin-gonic/gin"
)

func NewWebService(s interfaces.Service) interfaces.WebService {
	g := gin.Default()

	ws := &gin_webservice.GinWebService{
		Service: s,
		Engine:  g,
	}

	g.GET("/clients/connected", ws.GetConnectedClientsIds)
	g.POST("/clients/message", ws.SendMessageToClientsByIds)

	return ws
}
