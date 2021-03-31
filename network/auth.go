package network

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
)

func PostNewAuth(name, pub string) bool {
	m := map[string]string{"username": name, "pub": pub}
	asBytes, _ := json.Marshal(m)
	return DoPost("auth", asBytes) != ""
}

func CreateUserKey(c *gin.Context) {
	m := mapJsonBody(c)
	name := m["username"]
	pub := m["pub"]
	UniverseLock.Lock()
	defer UniverseLock.Unlock()
	if len(universes[uids[uidIndex]].UsernameKeys[name]) == 0 {
		universes[uids[uidIndex]].UsernameKeys[name] = []byte(pub)
		universes[uids[uidIndex]].UserCreatedAt[name] = time.Now().Unix()
		universes[uids[uidIndex]].UserCount++
		c.JSON(200, gin.H{"ok": true})
	} else {
		c.JSON(422, gin.H{"ok": false})
	}
}
