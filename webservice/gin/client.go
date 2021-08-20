package gin

import (
	"github.com/aidar-darmenov/message-delivery/model"
	"github.com/gin-gonic/gin"
)

func (ws *GinWebService) GetConnectedClientsIds(c *gin.Context) {
	c.JSON(200, ws.Service.GetClients().Ids)
}

func (ws *GinWebService) SendMessageToClientsByIds(c *gin.Context) {

	var message model.MessageToClients

	err := c.Bind(&message)
	if err != nil {
		c.JSON(400, err)
	}

	exp := ws.Service.SendMessageToClientsByIds(message)
	if exp != nil {
		c.JSON(exp.HTTPCode, exp.ErrorMessage)
		return
	}

	c.JSON(200, "ok")
}
