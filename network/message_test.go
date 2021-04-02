package network

import "testing"

func TestMessage(t *testing.T) {
	m := Message{}
	m.Text = "hi"

	m.Score()
	//if testJson != expected {
	//	t.Fatal()
	//}
}
