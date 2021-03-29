package main

import (
	"clt/cli"
	"clt/network"
	"clt/persist"
	"fmt"
	"math/rand"
	"os"
	"strings"
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

	if len(cli.UniverseIds) == 0 {
		ids := network.MakeUniverses()
		persist.SaveToFile("UNIVERSE_IDS", strings.Join(ids, ","))
	} else {
		network.MakeUniversesWithIds(cli.UniverseIds)
	}

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "start" {
		cli.EnsureParamPass("port", "host")
		c := make(chan bool, 1)
		go network.Start(c, cli.ArgMap["port"], cli.ArgMap["host"])
		<-c
	} else if command == "config" {
		network.NewUniverse()
	}
}
