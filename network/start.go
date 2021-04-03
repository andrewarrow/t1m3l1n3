package network

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

var uidIndex byte = 0
var uids []string = []string{}
var universes map[string]*Universe = map[string]*Universe{}
var maxUniverses byte
var globalInOut *InOut

func Start(c chan bool, uidIndexString, port, host string) {

	fmt.Println("starting...")

	if uidIndexString != "" {
		b, _ := strconv.Atoi(uidIndexString)
		uidIndex = byte(b)
	}

	r := gin.Default()
	r.GET("/timelines", ShowRecent)
	r.GET("/timelines/:username", ShowTimelines)
	r.POST("/timelines", CreateTimeline)
	r.POST("/timelines_as", CreateTimelineAs)
	r.POST("/timelines/notify", NotifyTimeline)
	r.POST("/follow/:username", ToggleFollowPost)
	r.GET("/universe", ShowUniverse)
	r.POST("/auth", CreateUserKey)
	r.GET("/idplease", IdPlease)
	r.GET("/taken", ShowUsers)
	r.POST("/suggest", Suggest)
	r.POST("/score", Score)
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
