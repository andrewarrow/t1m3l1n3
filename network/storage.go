package network

import (
	"fmt"
	"log"
	"time"
)

const size = 64

type Universe struct {
	Following []uint64
	Inboxes   [][]*Timeline
	Usernames map[string]byte
	UserCount byte
}

func hasBit(n uint64, pos byte) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func (u *Universe) BroadcastNewTimeline(t *Timeline) {
	log.Println("BroadcastNewTimeline")
	for i := byte(0); i < u.UserCount; i++ {
		log.Println("ShouldDeliverFrom", t.From, i)
		if u.ShouldDeliverFrom(t.From, i) {
			log.Println(" TRUE")
			u.Inboxes[i] = append([]*Timeline{t}, u.Inboxes[i]...)
		}
	}
	log.Println("END BroadcastNewTimeline")
}

func (u *Universe) ShouldDeliverFrom(username string, to byte) bool {
	log.Println("  ShouldDeliverFrom", username, to)
	from := u.UsernameToIndex(username) - 1
	log.Println("  ShouldDeliverFrom", from)
	return hasBit(u.Following[to], from)
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

	t := Timeline{}
	t.Text = "Welcome to CLT"
	t.From = "sysop"
	t.PostedAt = time.Now().Unix()
	welcome := []*Timeline{}
	for i := 0; i < 1; i++ {
		welcome = append(welcome, &t)
	}
	for i := 0; i < size; i++ {
		u.Following = append(u.Following, 0xFFFFFFFFFFFFFFFF)
		u.Inboxes = append(u.Inboxes, welcome)
	}

	fmt.Println(u.Following)
	fmt.Println(u.Inboxes)
	return &u
}
