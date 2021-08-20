package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ws *GinWebService) GetConnectedClientsIds(c *gin.Context) {

	fmt.Println(ws.Service.GetConnectedClientsIds())

	c.JSON(200, "ok")
}
