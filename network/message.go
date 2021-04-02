package network

import (
	"fmt"
	"strconv"
)

type Message struct {
	Text string
}

func AsciiByteToBase9(a string) byte {

	sum := byte(0)
	for i := range a {

		word := a[i : i+1]
		t, _ := strconv.Atoi(word)

		sum += byte(t)
	}
	strSum := fmt.Sprintf("%d", sum)
	if len(strSum) > 1 {
		return AsciiByteToBase9(strSum)
	}
	return sum

}

func (m *Message) Score() {
	for _, a := range m.Text {
		num := fmt.Sprintf("%d", a)
		base9 := AsciiByteToBase9(num)
		fmt.Printf("%d ", base9)
	}
	fmt.Printf("\n")
}
