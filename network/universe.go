package network

import (
	"github.com/gin-gonic/gin"
)

func ShowUniverse(c *gin.Context) {
	c.JSON(200, gin.H{"uids": uids})
}
