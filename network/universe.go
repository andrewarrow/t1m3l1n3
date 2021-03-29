package network

import (
	"clt/cli"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Universe struct {
	Following        []uint64
	Inboxes          map[byte][]*Timeline
	Usernames        map[string]byte
	UsernameKeys     map[string][]byte
	UsernameVerified map[string]string
	Profile          map[byte][]*Timeline
	UserCount        byte
	Id               string
}

func MakeUniverses() []string {
	for i := 0; i < size; i++ {
		u := NewUniverse()
		uids = append(uids, u.Id)
		universes[u.Id] = u
	}
	return uids
}
func MakeUniversesWithIds(ids []string) []string {
	for i := 0; i < size; i++ {
		u := NewUniverse()
		u.Id = ids[i]
		uids = append(uids, u.Id)
		universes[u.Id] = u
	}
	return uids
}
func ShowUniverse(c *gin.Context) {
	items := []map[string]interface{}{}
	for _, id := range uids {
		stats := universes[id].MakeStats()
		items = append(items, stats)
	}
	c.JSON(200, gin.H{"items": items})
}

func (u *Universe) MakeStats() map[string]interface{} {
	m := map[string]interface{}{}

	m["user_count"] = u.UserCount
	m["id"] = u.Id
	m["inboxes_count"] = len(u.Inboxes)
	inboxesStats := map[byte]int{}
	profileStats := map[byte]int{}
	for k, v := range u.Inboxes {
		inboxesStats[k] = len(v)
	}
	for k, v := range u.Profile {
		profileStats[k] = len(v)
	}
	m["inboxes"] = inboxesStats
	m["profile"] = profileStats

	return m
}
func (u *Universe) BroadcastNewTimeline(t *Timeline) bool {
	fromIndex := u.UsernameToIndex(t.From) - 1
	if fromIndex == 255 {
		// no more room
		return false
	}
	u.Profile[fromIndex] = append([]*Timeline{t}, u.Profile[fromIndex]...)
	for i := byte(0); i < u.UserCount; i++ {
		if u.ShouldDeliverFrom(fromIndex, i) {
			u.Inboxes[i] = append([]*Timeline{t}, u.Inboxes[i]...)
		}
	}
	return true
}

func (u *Universe) ToggleFollow(to, from string) string {
	toIndex := u.UsernameToIndex(to) - 1
	fromIndex := u.UsernameToIndex(from) - 1
	u.Following[toIndex] = uint64(Bits(u.Following[toIndex]) ^ LookupBit(fromIndex))
	return fmt.Sprintf("%b", u.Following[toIndex])
}
func (u *Universe) ShouldDeliverFrom(from, to byte) bool {
	log.Println("  ShouldDeliverFrom", from)
	return HasBits(Bits(u.Following[to]), LookupBit(from))
}

func (u *Universe) UsernameToIndex(username string) byte {
	log.Println("    UsernameToIndex", username)
	if u.Usernames[username] == 0 {
		if u.UserCount == size {
			log.Println("    SIZE")
			return 0
		}
		u.UserCount++
		u.Usernames[username] = u.UserCount
	}
	log.Println("    u.Usernames[username]", u.Usernames[username])
	return u.Usernames[username]
}

func NewUniverse() *Universe {
	u := Universe{}
	u.Id = cli.MakeUuid()
	u.Following = []uint64{}

	u.Usernames = map[string]byte{}
	u.UsernameToIndex("sysop")
	u.Profile = map[byte][]*Timeline{}
	u.Inboxes = map[byte][]*Timeline{}

	t := Timeline{}
	t.Text = "Welcome to CLT"
	t.From = "sysop"
	t.PostedAt = time.Now().Unix()
	u.Profile[0] = []*Timeline{&t}
	for i := 0; i < size; i++ {
		u.Following = append(u.Following, 0xFFFFFFFFFFFFFFFF)
	}
	for i := byte(0); i < u.UserCount; i++ {
		welcome := []*Timeline{&t}
		u.Inboxes[i] = welcome
	}

	fmt.Println(u.Following)
	fmt.Println(u.Inboxes)
	return &u
}
