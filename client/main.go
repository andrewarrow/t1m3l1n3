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
	fmt.Println("  client ls        # List recent timelines")
	fmt.Println("  client profile   # List recent timelines")
	fmt.Println("  client post      # Post new timeline with --text=hi")
	fmt.Println("  client keygen    # Generate public/private keys")
	fmt.Println("  client sign      # Sign a message")
	fmt.Println("  client auth      # Set your username --name=")
	fmt.Println("  client servers   # List the main list")
	fmt.Println("  client simulate  # Simulate traffic")
	fmt.Println("  client toggle    # Toggle follow --name=")
	fmt.Println("  client universe  # Display universe json")
	fmt.Println("  client idplease  # tell me what server/node i'm connected to")
	fmt.Println("  client taken     # taken usernames")
	fmt.Println("  client suggest   # suggest a new universe with less users")
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

	if persist.ReadFromFile("PUBLIC_KEY") == "" {
		KeyGen()
	}

	if command == "profile" {
		username := cli.Username
		if cli.ArgMap["name"] != "" {
			username = cli.ArgMap["name"]
		}
		s := network.DoGet(fmt.Sprintf("timelines/%s", username))
		//fmt.Println(s)
		network.DisplayProfileTimelines(s)
	} else if command == "taken" {
		s := network.DoGet(fmt.Sprintf("taken"))
		fmt.Println(s)
	} else if command == "toggle" {
		cli.EnsureParamPass("name")
		s := network.DoPost(fmt.Sprintf("follow/%s", cli.ArgMap["name"]), []byte{})
		fmt.Println(s)
	} else if command == "suggest" {
		info := network.SuggestNewPlaceToAuth()
		if info["new_place_avail"].(bool) {
			s := info["server"].(string)
			i := byte(info["index"].(float64))
			persist.SaveToFile("SERVER", s)
			persist.SaveToFile("INDEX", fmt.Sprintf("%d", i))
			fmt.Println("Ok", s, i)
		} else {
			fmt.Println("Sorry I have no idea.")
		}
	} else if command == "keygen" {
		cli.EnsureParamPass("overwrite")
		KeyGen()
	} else if command == "idplease" {
		s := persist.ReadFromFile("SERVER")
		i := persist.ReadFromFile("INDEX")
		fmt.Println(s, i)
	} else if command == "verify" {
		DoTestSignAndVerify()
	} else if command == "sign" {
		KeySign("test")
	} else if command == "ls" {
		s := network.DoGet(fmt.Sprintf("timelines"))
		//fmt.Println(s)
		network.DisplayInboxTimelines(s)
	} else if command == "universe" {
		s := network.DoGet(fmt.Sprintf("universe"))
		fmt.Println(s)
	} else if command == "servers" {
		s := network.DoGet("servers")
		fmt.Println(s)
	} else if command == "auth" {
		cli.EnsureParamPass("name")
		pub := persist.ReadFromFile("PUBLIC_KEY")
		auth := network.PostNewAuth(cli.ArgMap["name"], pub)
		if auth["user_created"].(bool) {
			s := auth["server"].(string)
			i := byte(auth["index"].(float64))
			persist.SaveToFile("USERNAME", cli.ArgMap["name"])
			persist.SaveToFile("SERVER", s)
			persist.SaveToFile("INDEX", fmt.Sprintf("%d", i))
			fmt.Println("Ok you are now:", cli.ArgMap["name"])
		} else {
			fmt.Println("This username already taken!")
		}
	} else if command == "simulate" {
		people := []string{"bob", "alice", "candy", "mike", "dave", "chris", "pam", "abigail", "emma", "luna",
			"logan", "owen", "liam", "sophia", "santiago", "joe", "dan", "mark", "charles", "kevin",
			"logan2", "owen2", "liam2", "sophia2", "santiago2", "joe2", "dan2", "mark2", "charles2", "kevin2",
			"logan3", "owen3", "liam3", "sophia3", "santiago3", "joe3", "dan3", "mark3", "charles3", "kevin3",
			"logan4", "owen4", "liam4", "sophia4", "santiago4", "joe4", "dan4", "mark4", "charles4", "kevin4",
			"logan5", "owen5", "liam5", "sophia5", "santiago5", "joe5", "dan5", "mark5", "charles5", "kevin5",
			"logan6", "owen6", "liam6", "sophia6"}
		words := []string{"hi", "there"}

		for _, person := range people {
			for _, word := range words {
				s := KeySign(word)
				network.PostNewTimeline(word, person, s)
				time.Sleep(time.Millisecond * 20)
			}
		}
	} else if command == "post" {
		cli.EnsureParamPass("text")
		s := KeySign(cli.ArgMap["text"])
		network.PostNewTimeline(cli.ArgMap["text"], s)
	}
}
