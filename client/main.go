package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  clt ls      # List recent timelines")
	fmt.Println("  clt post    # Post new timeline with --text=hi")
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "ls" {
	} else if command == "post" {
	}
}
