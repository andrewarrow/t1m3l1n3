package network

import "fmt"

func NewUniverse() {
	one64 := []uint64{}

	for i := 0; i < 64; i++ {
		// 18446744073709551615
		one64 = append(one64, 0xFFFFFFFFFFFFFFFF)
	}

	fmt.Println(one64)
}
