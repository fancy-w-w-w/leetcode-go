package build

import (
	"math/rand"
	"testing"
)

func BenchmarkPow1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pow2(rand.Intn(1000), rand.Intn(3))
	}
}

func BenchmarkPow3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pow3(rand.Intn(1000), rand.Intn(3))
	}
}
