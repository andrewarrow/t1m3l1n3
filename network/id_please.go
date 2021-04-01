package network

import (
	"github.com/gin-gonic/gin"
)

func IdPlease(c *gin.Context) {
	info := map[string]string{}
	c.JSON(200, gin.H{"inbox": info})
}
