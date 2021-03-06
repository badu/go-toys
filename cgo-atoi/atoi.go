package atoi

// #include <stdlib.h>
// int atoi2(void *buf){return atoi((char *)buf); }
import "C"
import "unsafe"

func CAtoi(buf []byte) (int, error) {
	// error is always nil
	return int(C.atoi2(unsafe.Pointer(&buf[0]))), nil
}

func Direct(b []byte) (uint16, error) {
	ret, c := uint16(0), byte(0)
	for _, c = range b {
		ret = ret*10 + uint16(c-'0')
	}
	return ret, nil
}

func DirectShift(b []byte) (uint16, error) {
	ret, c, k := uint16(0), byte(0), uint16(0)
	for _, c = range b {
		k = ret<<1 + uint16(c-'0')
		ret <<= 3
		ret += k
	}
	return ret, nil
}
