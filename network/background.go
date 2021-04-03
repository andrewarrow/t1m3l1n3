package network

import (
	"encoding/json"
	"log"
	"t1m3l1n3/persist"
	"time"
)

func BackgroundThread() {
	for {
		log.Println("Background Thread...")
		time.Sleep(time.Second * 60)
		UniverseLock.Lock()
		for id, u := range universes {
			asBytes, _ := json.Marshal(u.Marshal())
			persist.SaveToFile(id, string(asBytes))
		}
		UniverseLock.Unlock()
	}
}
