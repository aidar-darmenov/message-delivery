package gin

import "github.com/gin-gonic/gin"

func (ws *GinWebService) GetConnectedClients(c *gin.Context) {
	c.JSON(200, ws.Service.GetConnectedClients)
}
