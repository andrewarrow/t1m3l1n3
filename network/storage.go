package network

import (
	"fmt"
	"time"
)

const size = 64

type Universe struct {
	Base1   []uint64
	Inboxes [][]*Timeline
}

func NewUniverse() *Universe {
	u := Universe{}
	u.Base1 = []uint64{}

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
		u.Base1 = append(u.Base1, 0xFFFFFFFFFFFFFFFF)
		u.Inboxes = append(u.Inboxes, welcome)
	}

	fmt.Println(u.Base1)
	fmt.Println(u.Inboxes)
	return &u
}
