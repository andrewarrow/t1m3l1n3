package network

import (
	"github.com/gin-gonic/gin"
)

func ShowUniverse(c *gin.Context) {
	c.JSON(200, gin.H{"uid1": uid1, "uid2": uid2})
}
