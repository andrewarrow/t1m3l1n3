package network

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

var ServerLock sync.Mutex

var ServerMap map[string]string = map[string]string{"localhost:8080": "1", "localhost:8081": "1"}
var ServerList []string = []string{"localhost:8080", "localhost:8081"}

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
	ServerMap[name] = "1"
	ServerList = []string{}
	for k, _ := range ServerMap {
		ServerList = append(ServerList, fmt.Sprintf("%s", k))
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

type InOut struct {
	In  string `json:"in"`
	Out string `json:"out"`
}

func ParseInOut(s string) *InOut {
	var inOut InOut
	json.Unmarshal([]byte(s), &inOut)

	return &inOut
}
