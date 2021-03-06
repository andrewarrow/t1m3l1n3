package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"t1m3l1n3/cli"
	"t1m3l1n3/network"
	"t1m3l1n3/persist"
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

	uidIndexString := persist.ReadFromFile("UNIVERSE_INDEX")

	if len(cli.UniverseIds) == 0 {
		if cli.ArgMap["universes"] == "" {
			cli.ArgMap["universes"] = "6"
		}
		ids := network.MakeUniverses(cli.ArgMap["universes"])
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
		if cli.ArgMap["port"] == "" {
			cli.ArgMap["port"] = "8080"
		}
		if cli.ArgMap["host"] == "" {
			cli.ArgMap["host"] = "main"
		}
		c := make(chan bool, 1)
		go network.BackgroundThread()
		go network.Start(c, uidIndexString, cli.ArgMap["port"], cli.ArgMap["host"])
		<-c
	} else if command == "config" {
		network.NewUniverse()
	}
}
