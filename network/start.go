package network

import (
	"fmt"

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
	} else {
		s := `host=%s
port=%s
`
		payload := fmt.Sprintf(s, host, port)
		jsonString := DoPost("servers", []byte(payload))
		globalInOut = ParseInOut(jsonString)
		fmt.Println(globalInOut)
	}
	r.Run(":" + port)
}
