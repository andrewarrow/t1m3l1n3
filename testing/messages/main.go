package main

import (
	"clt/keys"
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
				s := network.DoPost("", "score", asBytes)
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

	bobPriv, bobPub := keys.KeyGen()
	network.PostNewAuth("bob", bobPub)
	suePriv, suePub := keys.KeyGen()
	network.PostNewAuth("sue", suePub)
	mikePriv, mikePub := keys.KeyGen()
	network.PostNewAuth("mike", mikePub)
	maryPriv, maryPub := keys.KeyGen()
	network.PostNewAuth("mary", maryPub)

	c := make(chan bool, 1)
	go PostAs(bobPriv, "bob", vg_more_m)
	go PostAs(suePriv, "sue", vg_more_f)
	go PostAs(mikePriv, "mike", lg_more_m)
	go PostAs(maryPriv, "mary", lg_more_f)
	<-c
}

func PostAs(priv, name string, messages []string) {
	for {
		fmt.Println(name)

		text := messages[rand.Intn(len(messages))]
		s := keys.KeySign(priv, text)
		network.PostNewTimeline(name, text, s)

		time.Sleep(time.Second * 2)
	}
}
