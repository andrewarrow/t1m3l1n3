package network

import (
	"encoding/json"
	"fmt"
	"strings"
	"t1m3l1n3/persist"
	"time"

	"github.com/gin-gonic/gin"
)

func Suggest(c *gin.Context) {
	c.JSON(200, gin.H{"new_place_avail": true, "server": "localhost:8080", "index": 1})
}

func SuggestNewPlaceToAuth() map[string]interface{} {
	jsonString := DoPost("", "", "suggest", []byte{})
	if jsonString == "" {
		return map[string]interface{}{}
	}
	var thing map[string]interface{}
	json.Unmarshal([]byte(jsonString), &thing)
	return thing
}

func PostNewAuth(name, pub string) map[string]interface{} {
	m := map[string]string{"username": name, "pub": pub}
	asBytes, _ := json.Marshal(m)
	jsonString := DoPost("", "", "auth", asBytes)
	if jsonString == "" {
		return map[string]interface{}{}
	}
	var thing map[string]interface{}
	json.Unmarshal([]byte(jsonString), &thing)
	return thing
}

func CreateUserKey(c *gin.Context) {
	m := mapJsonBody(c)
	name := strings.TrimSpace(m["username"])
	if len(name) < 2 || len(name) > 22 {
		c.JSON(422, gin.H{"error": "username between 2 and 22"})
		return
	}

	pub := m["pub"]
	UniverseLock.Lock()
	defer UniverseLock.Unlock()
	if universes[uids[uidIndex]].UserCount == maxUsersPerUniverse {
		uidIndex++
		persist.SaveToFile("UNIVERSE_INDEX", fmt.Sprintf("%d", uidIndex))
	}
	if len(universes[uids[uidIndex]].UsernameKeys[name]) == 0 {
		universes[uids[uidIndex]].UsernameKeys[name] = []byte(pub)
		universes[uids[uidIndex]].UserCreatedAt[name] = time.Now().Unix()
		universes[uids[uidIndex]].UserCount++
		c.JSON(200, gin.H{"user_created": true,
			"universe_id": uids[uidIndex],
			"server":      "localhost:8080",
			"index":       0})
	} else {
		c.JSON(200, gin.H{"user_created": false})
	}
}
