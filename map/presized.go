package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int, 1000000)
	_, hm := mapTypeAndValue(m)

	fmt.Printf("Elements | h.B | Buckets\n\n")

	fmt.Printf("%8d | %3d | %8d\n", hm.count, hm.B, 1<<hm.B)

	for i := 0; i < 1000000; i++ {
		m[i] = i * 3
	}

	fmt.Printf("%8d | %3d | %8d\n", hm.count, hm.B, 1<<hm.B)
}
