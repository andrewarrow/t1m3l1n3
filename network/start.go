package network

import (
	"fmt"
	"time"
)

func Start(c chan bool) {

	fmt.Println("starting...")

	for {
		time.Sleep(time.Second * 1)
	}

}
