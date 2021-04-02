package network

import "fmt"

type Message struct {
	Text string
}

func (m *Message) Score() {
	fmt.Println("Wefwe")
}
