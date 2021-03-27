package network

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

var ServerLock sync.Mutex

var ServerList []string = []string{"localhost:8080"}

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
	for _, s := range ServerList {
		if s == name {
			break
		}
	}
	index := -1
	for i, s := range ServerList {
		if s == name {
			index = i
			break
		}
	}
	if index == -1 {
		ServerList = append(ServerList, name)
	}
	if index > 0 {
		in = ServerList[index-1]
	}
	if index != len(ServerList)-1 {
		out = ServerList[index+1]
	} else if ServerList[0] != name {
		out = ServerList[0]
	}

	if globalInOut.Flavor == "main" && len(ServerList) > 1 {
		globalInOut.Out = ServerList[1]
		globalInOut.In = ServerList[len(ServerList)-1]
		if globalInOut.In == globalInOut.Out {
		}
		fmt.Println(globalInOut.Debug())
	}
	ServerLock.Unlock()
	//	TellOutAboutNewServer(&t, globalInOut.Out)

	c.JSON(200, gin.H{"in": in, "out": out, "name": name})
}

type InOut struct {
	In     string `json:"in"`
	Out    string `json:"out"`
	Name   string `name:"name"`
	Flavor string `name:"flavor"`
}

func (io *InOut) Debug() string {
	return fmt.Sprintf("In: %s, Out: %s, Name: %s |%s|", io.In, io.Out, io.Name, io.Flavor)
}

func ParseInOut(s string) *InOut {
	var inOut InOut
	json.Unmarshal([]byte(s), &inOut)

	return &inOut
}
