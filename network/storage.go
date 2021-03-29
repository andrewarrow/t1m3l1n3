package network

import (
	"fmt"
	"log"
	"time"
)

const size = 64

type Universe struct {
	Following []uint64
	Inboxes   map[byte][]*Timeline
	Usernames map[string]byte
	Profile   map[byte][]*Timeline
	UserCount byte
}

func (u *Universe) BroadcastNewTimeline(t *Timeline) {
	log.Println("BroadcastNewTimeline")
	fromIndex := u.UsernameToIndex(t.From) - 1
	u.Profile[fromIndex] = append([]*Timeline{t}, u.Profile[fromIndex]...)
	for i := byte(0); i < u.UserCount; i++ {
		log.Println("ShouldDeliverFrom", t.From, i)
		if u.ShouldDeliverFrom(fromIndex, i) {
			u.Inboxes[i] = append([]*Timeline{t}, u.Inboxes[i]...)
		}
	}
	log.Println("END BroadcastNewTimeline")
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
	if u.UserCount == size {
		log.Println("    SIZE")
		return 0
	}
	if u.Usernames[username] == 0 {
		u.UserCount++
		u.Usernames[username] = u.UserCount
	}
	log.Println("    u.Usernames[username]", u.Usernames[username])
	return u.Usernames[username]
}

func NewUniverse() *Universe {
	u := Universe{}
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
