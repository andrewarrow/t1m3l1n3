package network

import (
	"clt/cli"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var uid1 string
var uid2 string
var universes map[string]*Universe = map[string]*Universe{}
var globalInOut *InOut

func Start(c chan bool, port, host string) {

	uid1 = cli.MakeUuid()
	uid2 = cli.MakeUuid()
	universes[uid1] = NewUniverse()
	universes[uid2] = NewUniverse()
	fmt.Println("starting...")

	r := gin.Default()
	r.GET("/timelines", ShowInbox)
	r.GET("/timelines/:username", ShowTimelines)
	r.POST("/timelines", CreateTimeline)
	r.POST("/timelines/notify", NotifyTimeline)
	r.POST("/follow/:username", ToggleFollowPost)
	r.GET("/universe", ShowUniverse)
	if host == "main" {
		r.GET("/servers", ShowServers)
		r.POST("/servers", AddServer)
		globalInOut = &InOut{}
		globalInOut.Flavor = "main"
		globalInOut.Name = "localhost:8080"
	} else {
		s := `host=%s
port=%s
`
		payload := fmt.Sprintf(s, host, port)
		os.Setenv("CLT_HOST", "")
		jsonString := DoPost("servers", []byte(payload))
		globalInOut = ParseInOut(jsonString)
		globalInOut.Flavor = "other"
	}
	r.Run(":" + port)
}
