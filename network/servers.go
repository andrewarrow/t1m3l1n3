package network

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var ServerListLock sync.Mutex

var ServerList map[string]string = map[string]string{"neverssl.com": "80", "cyborg.st": "80"}

func ShowServers(c *gin.Context) {
	ServerListLock.Lock()
	c.JSON(200, gin.H{"servers": ServerList})
	ServerListLock.Unlock()
}
func AddServer(c *gin.Context) {
	m := mapBody(c)
	ServerListLock.Lock()
	ServerList[m["host"]] = m["port"]
	ServerListLock.Unlock()
	c.JSON(200, gin.H{"ok": true})
}
