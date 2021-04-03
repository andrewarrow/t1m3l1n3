package network

import "github.com/gin-gonic/gin"

func Score(c *gin.Context) {
	m := mapJsonBody(c)
	msg := Message{}
	msg.Text = m["text"]

	c.JSON(200, gin.H{"score": msg.Score()})
}
