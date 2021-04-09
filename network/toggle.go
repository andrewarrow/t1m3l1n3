package network

import (
	b64 "encoding/base64"
	"fmt"
	"t1m3l1n3/keys"

	"github.com/gin-gonic/gin"
)

func ToggleFollowPost(c *gin.Context) {
	m := mapJsonBody(c)
	sig := c.Request.Header["Username"][0]
	from := m["from"]
	uid := c.Request.Header["Universe"][0]
	to := m["to"]
	prefix := m["prefix"]
	UniverseLock.Lock()
	defer UniverseLock.Unlock()
	u := UniverseSearchByPrefix(prefix)
	if u == nil {
		c.JSON(401, gin.H{"": ""})
		return
	}
	universes[uid].ToggleFollow(sig, from, u, to)
	c.JSON(200, gin.H{"": ""})
}

func (u *Universe) ToggleFollow(sig, from string, other *Universe, to string) string {
	pub := u.UsernameKeys[from]
	sDec, _ := b64.StdEncoding.DecodeString(sig)

	if keys.VerifySig(pub, from, sDec) == false {
		return "fail"
	}
	fmt.Printf("ToggleFollow from %s to %s \n %s \n %s \n", from, to, u.Id, other.Id)
	if u.Block[other.Id] == nil {
		u.Block[other.Id] = map[string]map[string]bool{}
	}
	if u.Block[other.Id][from] == nil {
		u.Block[other.Id][from] = map[string]bool{}
	}
	u.Block[other.Id][from][to] = !u.Block[other.Id][from][to]

	return ""
}
