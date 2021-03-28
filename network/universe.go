package network

import (
	"clt/cli"

	"github.com/gin-gonic/gin"
)

func ShowUniverse(c *gin.Context) {
	c.JSON(200, gin.H{"server_id": cli.ServerId})
}
