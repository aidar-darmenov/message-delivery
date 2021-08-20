package gin

import (
	"github.com/aidar-darmenov/message-delivery/interfaces"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (ws *GinWebService) Start() {
	ws.Engine.Run(":" + strconv.Itoa(ws.Service.GetConfigParams().HttpPort))
}

type GinWebService struct {
	Service interfaces.Service
	Engine  *gin.Engine
}
