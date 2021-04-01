package network

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var uidIndex byte = 0
var uids []string = []string{}
var universes map[string]*Universe = map[string]*Universe{}
var maxUniverses byte
var globalInOut *InOut

func Start(c chan bool, port, host string) {

	fmt.Println("starting...")

	r := gin.Default()
	r.GET("/timelines", ShowInbox)
	r.GET("/timelines/:username", ShowTimelines)
	r.POST("/timelines", CreateTimeline)
	r.POST("/timelines/notify", NotifyTimeline)
	r.POST("/follow/:username", ToggleFollowPost)
	r.GET("/universe", ShowUniverse)
	r.POST("/auth", CreateUserKey)
	r.GET("/idplease", IdPlease)
	r.GET("/taken", ShowUsers)
	if host == "main" {
		r.GET("/servers", ShowServers)
		r.POST("/servers", AddServer)
		globalInOut = &InOut{}
		globalInOut.Flavor = "main"
		globalInOut.Name = "localhost:8080"
	} else {
		//TODO
	}
	r.Run(":" + port)
}
