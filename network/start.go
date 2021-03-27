package network

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Start(c chan bool) {

	fmt.Println("starting...")

	r := gin.Default()
	r.GET("/timelines", ShowTimelines)
	r.Run()
}
