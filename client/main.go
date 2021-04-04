package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"t1m3l1n3/cli"
	"t1m3l1n3/keys"
	"t1m3l1n3/network"
	"t1m3l1n3/persist"
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
	fmt.Println("  client change    # --server=x --index=y")
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
		priv, pub := keys.KeyGen()
		keys.KeyGenSave(priv, pub)
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
	} else if command == "change" {
		cli.EnsureParamPass("server", "index")
		persist.SaveToFile("SERVER", cli.ArgMap["server"])
		persist.SaveToFile("INDEX", cli.ArgMap["index"])
	} else if command == "toggle" {
		cli.EnsureParamPass("name")
		data := persist.ReadFromFile("PRIVATE_KEY")
		sig := keys.KeySign(data, cli.Username)
		tokens := strings.Split(cli.ArgMap["name"], "-")
		prefix := tokens[0]
		to := tokens[1]
		asBytes, _ := json.Marshal(map[string]string{"from": cli.Username,
			"prefix": prefix})
		uid := persist.ReadFromFile("UNIVERSE")
		s := network.DoPost(uid, sig, fmt.Sprintf("follow/%s", to), asBytes)
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
		priv, pub := keys.KeyGen()
		keys.KeyGenSave(priv, pub)
	} else if command == "idplease" {
		s := persist.ReadFromFile("SERVER")
		i := persist.ReadFromFile("INDEX")
		fmt.Println(s, i)
	} else if command == "verify" {
		keys.DoTestSignAndVerify()
	} else if command == "sign" {
		data := persist.ReadFromFile("PRIVATE_KEY")
		keys.KeySign(data, "test")
	} else if command == "ls" {
		s := network.DoGet(fmt.Sprintf("timelines"))
		//fmt.Println(s)
		network.DisplayRecentTimelines(s)
	} else if command == "universe" {
		s := network.DoGet(fmt.Sprintf("universe"))
		fmt.Println(s)
	} else if command == "servers" {
		s := network.DoGet("servers")
		fmt.Println(s)
	} else if command == "auth" {
		if cli.ArgMap["clear"] == "true" {
			persist.RemoveList(persist.AllFiles())
			return
		}
		cli.EnsureParamPass("name")
		pub := persist.ReadFromFile("PUBLIC_KEY")
		auth := network.PostNewAuth(cli.ArgMap["name"], pub)
		if auth["user_created"].(bool) {
			s := auth["server"].(string)
			uid := auth["universe_id"].(string)
			persist.SaveToFile("USERNAME", cli.ArgMap["name"])
			persist.SaveToFile("SERVER", s)
			persist.SaveToFile("UNIVERSE", uid)
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
				network.PostNewTimelineAs(word, person)
				time.Sleep(time.Millisecond * 20)
			}
		}
	} else if command == "post" {
		cli.EnsureParamPass("text")
		data := persist.ReadFromFile("PRIVATE_KEY")
		uid := persist.ReadFromFile("UNIVERSE")
		s := keys.KeySign(data, cli.ArgMap["text"])
		network.PostNewTimeline(uid, cli.Username, cli.ArgMap["text"], s)
	}
}
