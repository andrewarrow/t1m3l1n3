package network

import (
	"fmt"
	"testing"
)

func TestMessage(t *testing.T) {
	m := Message{}
	m.Text = "hello this message will be scored."
	ms := m.Score()
	fmt.Println(ms.Debug())
	// 5 2 9 9 3 5 8 5 6 7 5 1 2 7 7 7 4 2 5 2 6 9 9 5 8 2 5 7 9 3 6 2 1 1
	// f m 9 9 3 f f f 6 f f m m f f f m m f m 6 9 9 f f m f f 9 3 6 m m m
	// f m 9 9 M f f f F f f m m f f f m m f m F 9 9 f f m f f 9 M F m m m
	// . m - - m . . . . . . m m . . . m m . m . - - . . m . . - m . m m m
	m.Text = "this is an entirely different kind of message."
	ms = m.Score()
	fmt.Println(ms.Debug())
	// fffffffffmfmmfffm.mfmfmmmfmmffffmmfmmfmmfffmmm
	m.Text = "Is there any reason to believe that this actually happened? Conveniently, it was published on April 1st. The story itself would be a great April Fools' prank. :)"
	// fffffffffmfmmfffm.mfmfmmmfmmffffmmfmmfmmfffmmm

	ms = m.Score()
	fmt.Println(ms.Debug())
	m.Text = "Respectfully, I disagree. The decision is much more personal and nuanced. There can be a myriad of reasons to change your name, including taking your partner's name if you get married. It could be an aesthetic choice or due to family estrangement."
	ms = m.Score()
	fmt.Println(ms.Debug())
	//if testJson != expected {
	//	t.Fatal()
	//}
}
