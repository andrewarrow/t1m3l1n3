package main

import (
	"clt/cli"
	"clt/network"
	"clt/persist"
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
	persist.Init()
	cli.ReadInGlobalVars()

	if cli.ServerId == "" {
		cli.ServerId = cli.MakeUuid()
		persist.SaveToFile("SERVER_ID", cli.ServerId)
	}

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "start" {
		cli.EnsureParamPass("port")
		c := make(chan bool, 1)
		go network.Start(c, cli.ArgMap["port"], cli.ArgMap["level"])
		<-c
	} else if command == "config" {
	}
}
