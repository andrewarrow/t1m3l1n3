package network

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Start(c chan bool, port, level string) {

	fmt.Println("starting...")

	r := gin.Default()
	r.GET("/timelines", ShowTimelines)
	r.POST("/timelines", CreateTimeline)
	if level == "main" {
		r.GET("/servers", ShowServers)
	}
	r.Run(":" + port)
}
