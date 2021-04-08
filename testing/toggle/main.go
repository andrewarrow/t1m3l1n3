package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"t1m3l1n3/keys"
	"t1m3l1n3/network"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	bobPriv, bobPub := keys.KeyGen()
	authJson := network.PostNewAuth("bob", bobPub)
	bobUid := authJson["universe_id"].(string)

	suePriv, suePub := keys.KeyGen()
	authJson = network.PostNewAuth("sue", suePub)
	sueUid := authJson["universe_id"].(string)

	text := "hi"
	sig := keys.KeySign(bobPriv, text)
	network.PostNewTimeline(bobUid, "bob", text, sig)

	sig = keys.KeySign(suePriv, text)
	network.PostNewTimeline(sueUid, "sue", text, sig)

	s := network.DoGet(bobUid, "bob", fmt.Sprintf("timelines"))
	fmt.Println(s)
	s = network.DoGet(sueUid, "sue", fmt.Sprintf("timelines"))
	fmt.Println(s)

	tokens := strings.Split(sueUid, "-")
	prefix := tokens[1]

	sig = keys.KeySign(bobPriv, "bob")
	asBytes, _ := json.Marshal(map[string]string{"from": "bob",
		"to":     "sue",
		"prefix": prefix})
	s = network.DoPost(bobUid, sig, fmt.Sprintf("toggle"), asBytes)
	fmt.Println(s)
	//network.DisplayRecentTimelines(uid, username, s)
	s = network.DoGet(bobUid, "bob", fmt.Sprintf("timelines"))
	fmt.Println(s)
	s = network.DoGet(sueUid, "sue", fmt.Sprintf("timelines"))
	fmt.Println(s)
}
