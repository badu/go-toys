package p95

import (
	"math"
	"sort"
	"testing"
)

func find95usual(s []float64) float64 {
	if len(s) <= 1 {
		return s[0]
	}
	discard := float64(len(s)-1) * .05
	x := make([]float64, len(s))
	copy(x, s)
	sort.Sort(sort.Reverse(sort.Float64Slice(x)))

	pos := int(discard)         // position on which p95 begins
	y := discard - float64(pos) // min is 0, never 1
	return x[pos]*(1-y) + x[pos+1]*y
}

func find95short(s []float64) float64 {
	if len(s) <= 1 {
		return s[0] // panic for empty array, as designed
	}

	discard := float64(len(s)-1) * .05
	pos := int(discard)

	x := make([]float64, (pos+1)*2+pos/5) // a cut with 5% of elements and padding for rest of them, pos/5 helps with extra coupls spots on big arrays
	copy(x, s[0:len(x)])
	sort.Sort(sort.Reverse(sort.Float64Slice(x)))

	if len(s) > len(x) {
		fill, edge := pos+1, x[pos+1]
		for _, v := range s[len(x):] {
			if v > edge {
				x[fill] = v
				fill++
				if fill == len(x) {
					sort.Sort(sort.Reverse(sort.Float64Slice(x)))
					fill, edge = pos+1, x[pos+1]
				}
			}
		}
		if fill > pos+1 {
			sort.Sort(sort.Reverse(sort.Float64Slice(x)))
		}
	}

	y := discard - float64(pos)
	return x[pos]*(1-y) + x[pos+1]*y
}

func TestP95Correctness(t *testing.T) {
	for i := 257; i < 258; i++ {
		pu := find95usual(TESTSET[:i:i])
		pf := find95short(TESTSET[:i:i])
		if math.Abs(pu-pf) > 0.00001 {
			t.Errorf("Incorrect p95, %.4f != %.4f  N %d diff %.5f", pu, pf, i, math.Abs(pu-pf))
		}
	}
}

func BenchmarkP95Usual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = find95usual(TESTSET[:NUM:NUM])
	}
}

func BenchmarkP95Short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = find95short(TESTSET[:NUM:NUM])
	}
}
