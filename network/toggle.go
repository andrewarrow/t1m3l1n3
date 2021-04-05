package network

import (
	b64 "encoding/base64"
	"t1m3l1n3/keys"

	"github.com/gin-gonic/gin"
)

func ToggleFollowPost(c *gin.Context) {
	m := mapJsonBody(c)
	sig := c.Request.Header["Username"]
	uid := c.Request.Header["Universe"]
	to := c.Param("username")
	prefix := m["prefix"]
	UniverseLock.Lock()
	defer UniverseLock.Unlock()
	u := UniverseSearchByPrefix(prefix)
	if u == nil {
		c.JSON(401, gin.H{"": ""})
		return
	}
	universes[uid[0]].ToggleFollow(sig[0], m["from"], u, to)
	c.JSON(200, gin.H{"": ""})
}

func (u *Universe) ToggleFollow(sig, from string, other *Universe, to string) string {
	pub := u.UsernameKeys[from]
	sDec, _ := b64.StdEncoding.DecodeString(sig)

	if keys.VerifySig(pub, from, sDec) == false {
		return "fail"
	}
	if u.Block[other.Id] == nil {
		u.Block[other.Id] = &BlockThing{}
	}
	if u.Block[other.Id].Thing == nil {
		u.Block[other.Id].Thing = map[string]*OtherBlockThing{}
	}
	if u.Block[other.Id].Thing[from] == nil {
		u.Block[other.Id].Thing[from].Thing = map[string]*LastBlockThing{}
	}
	foo := u.Block[other.Id]
	if foo.Thing[from].Thing[to] == nil {
		foo.Thing[from].Thing[to] = &LastBlockThing{}
	}
	foo.Thing[from].Thing[to].Thing = !foo.Thing[from].Thing[to].Thing
	return ""
}
