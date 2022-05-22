package main

import (
	"math/rand"
	"testing"
	"time"
)

const size = 1000000

var ints = make([]int, 0, size)

func init() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		ints = append(ints, rand.Intn(size))
	}
}

func BenchmarkIntPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := IntPairs(ints)
		b.Logf("n is %d", n)
	}
}
