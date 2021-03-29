package network

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

var ServerLock sync.Mutex

var ServerLink map[string]string = map[string]string{"localhost:8080": ""}
var ServerList []string = []string{"localhost:8080"}

func ShowServers(c *gin.Context) {
	ServerLock.Lock()
	c.JSON(200, gin.H{"servers": ServerList})
	ServerLock.Unlock()
}
func AddServer(c *gin.Context) {
	m := mapJsonBody(c)
	name := fmt.Sprintf("%s:%s", m["host"], m["port"])
	//in := ""
	//out := ""
	ServerLock.Lock()
	index := -1
	for i, s := range ServerList {
		if s == name {
			index = i
			break
		}
	}
	/*
	   s1: s2
	   s2: s1

	   s1: s2
	   s2: s3
	   s3: s1
	*/
	if index == -1 {
		previous := ServerList[len(ServerList)-1]
		ServerList = append(ServerList, name)
		latest := ServerList[len(ServerList)-1]
		ServerLink[latest] = globalInOut.Name
		ServerLink[previous] = name
	}

	fmt.Println(ServerLink)
	ServerLock.Unlock()
	//	TellOutAboutNewServer(&t, globalInOut.Out)
	// s3 -> s1       // s1 tells s2 your new out is s3

	c.JSON(200, gin.H{"out": globalInOut.Name})
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
