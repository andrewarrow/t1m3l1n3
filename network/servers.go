package network

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

var ServerLock sync.Mutex

var ServerMap map[string]string = map[string]string{"neverssl.com": "80", "cyborg.st": "80"}
var ServerList []string = []string{"neverssl.com:80", "cyborg.st:80"}

func ShowServers(c *gin.Context) {
	ServerLock.Lock()
	c.JSON(200, gin.H{"servers": ServerList})
	ServerLock.Unlock()
}
func AddServer(c *gin.Context) {
	m := mapBody(c)
	ServerLock.Lock()
	ServerMap[m["host"]] = m["port"]
	ServerList = []string{}
	for k, v := range ServerMap {
		ServerList = append(ServerList, fmt.Sprintf("%s:%s", k, v))
	}
	ServerLock.Unlock()
	c.JSON(200, gin.H{"ok": true})
}
