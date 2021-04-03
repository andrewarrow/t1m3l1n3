package network

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"t1m3l1n3/cli"
	"t1m3l1n3/persist"

	"github.com/gin-gonic/gin"
)

type Universe struct {
	Following     []uint64
	Recent        []*Timeline
	Inboxes       map[byte][]*Timeline
	Usernames     map[string]byte
	UserCreatedAt map[string]int64
	UsernameKeys  map[string][]byte
	Profile       map[byte][]*Timeline
	UserCount     byte
	Id            string
	UpPeers       []string
	DownPeers     []string
}

func MakeUniverses(s string) []string {
	max, _ := strconv.Atoi(s)
	for i := 0; i < max; i++ {
		u := NewUniverse()
		uids = append(uids, u.Id)
		universes[u.Id] = u
	}
	root := universes[uids[0]]
	DownPeers := []string{}
	for i := 1; i < max; i++ {
		universes[uids[i]].UpPeers = []string{"localhost:8080,0"}
		DownPeers = append(DownPeers, fmt.Sprintf("localhost:8080,%d", i))
	}
	root.UpPeers = []string{}
	root.DownPeers = DownPeers

	return uids
}
func MakeUniversesWithIds(ids []string) []string {
	for i := 0; i < len(ids); i++ {
		u := NewUniverse()
		u.Id = ids[i]
		jsonString := persist.ReadFromFile(u.Id)
		if jsonString != "" {
			var m map[string]interface{}
			json.Unmarshal([]byte(jsonString), &m)
			f := m["following"].([]interface{})
			u.Following = []uint64{}
			for _, v := range f {
				val, _ := strconv.ParseUint(v.(string), 10, 64)
				u.Following = append(u.Following, val)
			}
			u.UserCount = byte(m["user_count"].(float64))
			other := m["usernames"].(map[string]interface{})
			for k, v := range other {
				u.Usernames[k] = byte(v.(float64))
			}
			other = m["user_created_at"].(map[string]interface{})
			for k, v := range other {
				u.UserCreatedAt[k] = int64(v.(float64))
			}
			other = m["username_keys"].(map[string]interface{})
			for k, v := range other {
				sDec, _ := b64.StdEncoding.DecodeString(v.(string))
				u.UsernameKeys[k] = sDec
			}
			f = m["up_peers"].([]interface{})
			for _, v := range f {
				u.UpPeers = append(u.UpPeers, v.(string))
			}
			f = m["down_peers"].([]interface{})
			for _, v := range f {
				u.DownPeers = append(u.DownPeers, v.(string))
			}
		}
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

func (u *Universe) Marshal() map[string]interface{} {
	m := map[string]interface{}{}
	FollowingPayload := []string{}
	UsernameKeysPayload := map[string]string{}

	for _, val := range u.Following {
		FollowingPayload = append(FollowingPayload, fmt.Sprintf("%d", val))
	}
	for k, val := range u.UsernameKeys {
		enc := b64.StdEncoding.EncodeToString(val)
		UsernameKeysPayload[k] = enc
	}
	m["user_count"] = u.UserCount
	m["id"] = u.Id
	m["following"] = FollowingPayload
	m["usernames"] = u.Usernames
	m["username_keys"] = UsernameKeysPayload
	m["user_created_at"] = u.UserCreatedAt
	m["up_peers"] = u.UpPeers
	m["down_peers"] = u.DownPeers

	return m
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
	u.Recent = append([]*Timeline{t}, u.Recent...)
	if len(u.Recent) > 100 {
		u.Recent = u.Recent[0:100]
	}
	u.Profile[fromIndex] = append([]*Timeline{t}, u.Profile[fromIndex]...)
	for i := byte(0); i < maxUsersPerUniverse; i++ {
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
		if u.UserCount == maxUsersPerUniverse {
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

	u.Recent = []*Timeline{}
	u.UpPeers = []string{}
	u.DownPeers = []string{}
	u.Usernames = map[string]byte{}
	u.UsernameKeys = map[string][]byte{}
	u.UserCreatedAt = map[string]int64{}
	u.Profile = map[byte][]*Timeline{}
	u.Inboxes = map[byte][]*Timeline{}

	for i := 0; i < maxUsersPerUniverse; i++ {
		u.Following = append(u.Following, 0xFFFFFFFFFFFFFFFF)
	}

	return &u
}
