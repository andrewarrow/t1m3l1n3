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
	fmt.Println("  clt ls       # List recent timelines")
	fmt.Println("  clt post     # Post new timeline with --text=hi")
	fmt.Println("  clt auth     # Set your username --name=")
	fmt.Println("  clt servers  # List the main list")
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	persist.Init()
	cli.ReadInGlobalVars()

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "ls" {
		s := network.DoGet("timelines")
		network.DisplayTimelines(s)
	} else if command == "servers" {
		s := network.DoGet("servers")
		fmt.Println(s)
	} else if command == "auth" {
		persist.SaveToFile("USERNAME", cli.ArgMap["name"])
	} else if command == "post" {
		s := `text=%s
username=%s
`
		payload := fmt.Sprintf(s, cli.ArgMap["text"], cli.Username)
		network.DoPost("timelines", []byte(payload))
	}
}
