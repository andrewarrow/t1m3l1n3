package main

import (
	"clt/network"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  clt start   # Listen on port --port=x")
	fmt.Println("  clt config  # Print config")
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "start" {
		c := make(chan bool, 1)
		go network.Start(c)
		<-c
	} else if command == "config" {
	}
}
