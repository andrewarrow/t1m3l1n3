package network

import (
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTimelineAs(c *gin.Context) {
	m := mapJsonBody(c)
	t := Timeline{}
	t.Text = m["text"]
	t.From = m["username"]

	t.PostedAt = time.Now().Unix()
	t.Origin = globalInOut.Name

	if t.AddToUniverse("") == true {
		c.JSON(200, gin.H{"ok": true})
		return
	}
	c.JSON(422, gin.H{"ok": false})
}
