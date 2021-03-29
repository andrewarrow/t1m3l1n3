package network

import (
	"clt/persist"
	"encoding/json"
	"log"
	"time"
)

func BackgroundThread() {
	for {
		log.Println("Background Thread...")
		time.Sleep(time.Second * 6)
		for id, u := range universes {
			asBytes, _ := json.Marshal(u.Marshal())
			persist.SaveToFile(id, string(asBytes))
		}
	}
}
