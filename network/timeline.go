package network

import "github.com/gin-gonic/gin"

func ShowTimelines(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
