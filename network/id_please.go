package network

import (
	"github.com/gin-gonic/gin"
)

func IdPlease(c *gin.Context) {
	info := map[string]string{}
	c.JSON(200, gin.H{"inbox": info})
}

func ShowUsers(c *gin.Context) {
	info := []map[string]interface{}{}
	UniverseLock.Lock()
	for username, ts := range universes[uids[uidIndex]].UserCreatedAt {
		m := map[string]interface{}{}
		m["username"] = username
		m["ts"] = ts
		info = append(info, m)
	}
	UniverseLock.Unlock()
	c.JSON(200, gin.H{"users": info})
}
