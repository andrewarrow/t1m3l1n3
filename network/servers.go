package network

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var ServerListLock sync.Mutex

var ServerList []string = []string{"http://neverssl.com:80", "http://cyborg.st:80"}

func ShowServers(c *gin.Context) {
	ServerListLock.Lock()
	c.JSON(200, gin.H{"servers": ServerList})
	ServerListLock.Unlock()
}
