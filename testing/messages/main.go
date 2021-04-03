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
	vg_more_m := []string{}
	vg_more_f := []string{}
	lg_more_m := []string{}
	lg_more_f := []string{}
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
				lr := int(score["longest_run"].(float64))
				fmt.Printf("%s %.2f %.2f %.2f, %d %d\n\n", lf, mp, fp, gp, sw, lr)

				if gp > 0.10 && mp > fp {
					vg_more_m = append(vg_more_m, thing)
				} else if gp > 0.10 && mp < fp {
					vg_more_f = append(vg_more_f, thing)
				} else if gp <= 0.10 && mp > fp {
					lg_more_m = append(lg_more_m, thing)
				} else if gp <= 0.10 && mp < fp {
					lg_more_f = append(lg_more_f, thing)
				}

				buff = []string{}
			}
		}
	}

	c := make(chan bool, 1)
	go PostAs("bob", vg_more_m)
	go PostAs("sue", vg_more_f)
	go PostAs("mike", lg_more_m)
	go PostAs("mary", lg_more_f)
	<-c
}

func PostAs(name string, messages []string) {
	for {
		fmt.Println(name)
		time.Sleep(time.Second * 10)
	}
}
