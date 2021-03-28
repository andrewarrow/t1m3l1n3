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
	Profile   map[byte][]*Timeline
	UserCount byte
}

type Bits uint64

const (
	F0 Bits = 1 << iota
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	F13
	F14
	F15
	F16
	F17
	F18
	F19
	F20
	F21
	F22
	F23
	F24
	F25
	F26
	F27
	F28
	F29
	F30
	F31
	F32
	F33
	F34
	F35
	F36
	F37
	F38
	F39
	F40
	F41
	F42
	F43
	F44
	F45
	F46
	F47
	F48
	F49
	F50
	F51
	F52
	F53
	F54
	F55
	F56
	F57
	F58
	F59
	F60
	F61
	F62
	F63
)

func Set(b, flag Bits) Bits    { return b | flag }
func Clear(b, flag Bits) Bits  { return b &^ flag }
func Toggle(b, flag Bits) Bits { return b ^ flag }
func Has(b, flag Bits) bool    { return b&flag != 0 }

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

func (u *Universe) ToggleFollow(from, to string) string {
	toIndex := u.UsernameToIndex(to) - 1
	/*
		fromIndex := u.UsernameToIndex(from) - 1
			if Has(u.Following[toIndex], fromIndex) {
				u.Following[toIndex] &= clearBit(u.Following[toIndex], fromIndex)
			} else {
				u.Following[toIndex] &= setBit(u.Following[toIndex], fromIndex)
			}
	*/
	return fmt.Sprintf("%b", u.Following[toIndex])
}
func (u *Universe) ShouldDeliverFrom(from, to byte) bool {
	log.Println("  ShouldDeliverFrom", from)
	//return hasBit(u.Following[to], from)
	return true
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

	t := Timeline{}
	t.Text = "Welcome to CLT"
	t.From = "sysop"
	t.PostedAt = time.Now().Unix()
	u.Profile[0] = []*Timeline{&t}
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
