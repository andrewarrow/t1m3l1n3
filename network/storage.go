package network

import (
	"fmt"
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
	for i := byte(0); i < u.UserCount; i++ {
		if u.ShouldDeliverFrom(t.From, i+1) {
			u.Inboxes[i] = append([]*Timeline{t}, u.Inboxes[i]...)
		}
	}
}

func (u *Universe) ShouldDeliverFrom(username string, to byte) bool {
	from := u.UsernameToIndex(username)
	return hasBit(u.Following[to], from)
}

func (u *Universe) UsernameToIndex(username string) byte {
	if u.UserCount == size {
		return 0
	}
	if u.Usernames[username] == 0 {
		u.UserCount++
		u.Usernames[username] = u.UserCount
	}
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
		// 18446744073709551615
		u.Following = append(u.Following, 0xFFFFFFFFFFFFFFFF)
		u.Inboxes = append(u.Inboxes, welcome)
	}

	fmt.Println(u.Following)
	fmt.Println(u.Inboxes)
	return &u
}
