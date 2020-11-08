package main

import (
	"fmt"
	"unsafe"
)

func showSpread(m interface{}) {
	// dataOffset is where the cell data begins in a bmap
	const dataOffset = unsafe.Offsetof(struct {
		tophash [bucketCnt]uint8
		cells   int64
	}{}.cells)

	t, h := mapTypeAndValue(m)

	fmt.Printf("Overflow buckets: %d", h.noverflow)

	numBuckets := 1 << h.B

	for r := 0; r < numBuckets*bucketCnt; r++ {
		bucketIndex := r / bucketCnt
		cellIndex := r % bucketCnt

		if cellIndex == 0 {
			fmt.Printf("\nBucket %3d:", bucketIndex)
		}

		// lookup cell
		b := (*bmap)(add(h.buckets, uintptr(bucketIndex)*uintptr(t.bucketsize)))
		if b.tophash[cellIndex] == 0 {
			// cell is empty
			continue
		}

		k := add(unsafe.Pointer(b), dataOffset+uintptr(cellIndex)*uintptr(t.keysize))

		ei := emptyInterface{
			_type: unsafe.Pointer(t.key),
			value: k,
		}
		key := *(*interface{})(unsafe.Pointer(&ei))
		fmt.Printf(" %3d", key.(int))
	}
	fmt.Printf("\n\n")
}

func main() {
	m := make(map[int]int)

	for i := 0; i < 50; i++ {
		m[i] = i * 3
	}

	showSpread(m)

	m = make(map[int]int)

	for i := 0; i < 8; i++ {
		m[i] = i * 3
	}

	showSpread(m)
}
