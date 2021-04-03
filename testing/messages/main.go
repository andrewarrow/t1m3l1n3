package main

import (
	"clt/network"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("messages")
	f, _ := ioutil.ReadFile("../latin.html")
	buff := []string{}
	for _, line := range strings.Split(string(f), "\n") {
		for _, word := range strings.Split(line, " ") {
			buff = append(buff, word)
			thing := strings.Join(buff, "")
			if len(thing) >= 20+rand.Intn(190) {
				fmt.Println(len(thing))
				m := map[string]string{}
				m["text"] = thing
				asBytes, _ := json.Marshal(m)
				s := network.DoPost("score", asBytes)
				fmt.Println(s)
				buff = []string{}
			}
		}
	}
}
