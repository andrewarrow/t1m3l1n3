package network

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var globalInOut *InOut

func Start(c chan bool, port, host string) {

	fmt.Println("starting...")

	r := gin.Default()
	r.GET("/timelines", ShowTimelines)
	r.POST("/timelines", CreateTimeline)
	r.POST("/timelines/notify", NotifyTimeline)
	if host == "main" {
		r.GET("/servers", ShowServers)
		r.POST("/servers", AddServer)
		globalInOut = &InOut{}
		globalInOut.Flavor = "main"
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
	fmt.Println(globalInOut.Debug())
	r.Run(":" + port)
}
