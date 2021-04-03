package network

import (
	"fmt"
	"strconv"
	"strings"
)

type Message struct {
	Text string `json:"text"`
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

type MessageScore struct {
	M             *Message
	PercentM      float64 `json:"percent_m"`
	PercentF      float64 `json:"percent_f"`
	PercentG      float64 `json:"percent_g"`
	LongestRun    int     `json:"longest_run"`
	LongestFlavor string  `json:"longest_flavor"`
	Switches      int     `json:"switches"`
	Joined        string  `json:"joined"`
}

func (m *Message) Score() *MessageScore {
	ms := MessageScore{}
	ms.M = m
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
	joined := strings.Join(buff, "")
	sumM := 0
	sumF := 0
	sumG := 0
	Switches := 0
	LongestRun := 0
	MaxLongestRun := 0
	Prev := ""
	MaxLongestRunFlavor := ""
	for i := range joined {
		c := joined[i : i+1]
		if c != Prev {
			Switches++
			if LongestRun > MaxLongestRun {
				MaxLongestRun = LongestRun
				MaxLongestRunFlavor = Prev
			}
			LongestRun = 0
		} else {
			LongestRun++
		}
		if c == "m" {
			sumM++
		} else if c == "f" {
			sumF++
		} else if c == "." {
			sumG++
		}
		Prev = c
	}
	ms.PercentM = float64(sumM) / float64(len(joined))
	ms.PercentF = float64(sumF) / float64(len(joined))
	ms.PercentG = float64(sumG) / float64(len(joined))
	ms.LongestRun = MaxLongestRun
	ms.LongestFlavor = MaxLongestRunFlavor
	ms.Joined = joined
	ms.Switches = Switches
	//fmt.Printf("%s\n%+v\n", joined, ms)
	return &ms
}

func (ms *MessageScore) Debug() string {
	buff := []string{}
	buff = append(buff, fmt.Sprintf("%s\n", ms.M.Text))
	buff = append(buff, fmt.Sprintf("%s\n", ms.Joined))
	buff = append(buff, fmt.Sprintf("%f %f %f\n", ms.PercentM, ms.PercentF, ms.PercentG))
	buff = append(buff, fmt.Sprintf("%s %d\n", ms.LongestFlavor, ms.LongestRun))
	buff = append(buff, fmt.Sprintf("%d\n\n", ms.Switches))
	return strings.Join(buff, "")
}
