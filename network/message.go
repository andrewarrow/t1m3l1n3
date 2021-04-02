package network

import (
	"fmt"
	"strconv"
	"strings"
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
	buff := []string{}
	for _, a := range m.Text {
		num := fmt.Sprintf("%d", a)
		base9 := fmt.Sprintf("%d", AsciiByteToBase9(num))
		if base9 == "5" || base9 == "7" || base9 == "8" || base9 == "6" {
			buff = append(buff, "f")
		} else if base9 == "3" || base9 == "1" || base9 == "2" || base9 == "4" {
			buff = append(buff, "m")
		} else if base9 == "9" {
			buff = append(buff, ".")
		}
	}
	fmt.Printf("%s\n", strings.Join(buff, ""))
}
