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
	name := fmt.Sprintf("%s:%s", m["host"], m["port"])
	in := ""
	out := ""
	ServerLock.Lock()
	ServerMap[m["host"]] = m["port"]
	ServerList = []string{}
	for k, v := range ServerMap {
		ServerList = append(ServerList, fmt.Sprintf("%s:%s", k, v))
	}
	index := 0
	for i, s := range ServerList {
		if s == name {
			index = i
			break
		}
	}
	if index > 0 {
		in = ServerList[index-1]
	}
	if index != len(ServerList)-1 {
		out = ServerList[index+1]
	} else if ServerList[0] != name {
		out = ServerList[0]
	}
	ServerLock.Unlock()
	c.JSON(200, gin.H{"in": in, "out": out})
}
