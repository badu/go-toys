package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)
	_, hm := mapTypeAndValue(m)

	fmt.Printf("Elements | h.B | Buckets\n\n")

	for i := 0; i < 100000; i++ {
		m[i] = i * 3
	}

	for i := 0; i < 100000; i++ {
		delete(m, i)
	}

	fmt.Printf("%8d | %3d | %8d\n", hm.count, hm.B, 1<<hm.B)
}
