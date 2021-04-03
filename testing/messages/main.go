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
	stats := map[int]int{}
	stats2 := map[int]int{}
	for _, line := range strings.Split(string(f), "\n") {
		for _, word := range strings.Split(line, " ") {
			buff = append(buff, word)
			thing := strings.Join(buff, "")
			if len(thing) >= 20+rand.Intn(190) {
				m := map[string]string{}
				m["text"] = thing
				asBytes, _ := json.Marshal(m)
				s := network.DoPost("score", asBytes)
				var mm map[string]interface{}
				json.Unmarshal([]byte(s), &mm)
				score := mm["score"].(map[string]interface{})
				lf := score["longest_flavor"].(string)
				mp := score["percent_m"].(float64)
				fp := score["percent_f"].(float64)
				gp := score["percent_g"].(float64)
				sw := int(score["switches"].(float64))
				stats[sw]++
				lr := int(score["longest_run"].(float64))
				stats2[lr]++
				fmt.Printf("%s %.2f %.2f %.2f, %d %d\n\n", lf, mp, fp, gp, sw, lr)

				buff = []string{}
			}
		}
	}
	fmt.Println(stats)
	fmt.Println(stats2)
}
