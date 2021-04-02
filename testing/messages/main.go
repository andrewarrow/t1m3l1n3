package main

import (
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
			if len(strings.Join(buff, "")) >= 20+rand.Intn(190) {
				fmt.Println(len(strings.Join(buff, "")))
				buff = []string{}
			}
		}
	}
}
