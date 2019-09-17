package absmax

import (
	"math"
	"testing"
)

func BenchmarkABSInt64(b *testing.B) {
	for i := -b.N; i <= b.N; i++ {
		AbsInt64(int64(i))
	}
}

func BenchmarkGoABSInt64(b *testing.B) {
	for i := -b.N; i <= b.N; i++ {
		math.Abs(float64(i))
	}
}
