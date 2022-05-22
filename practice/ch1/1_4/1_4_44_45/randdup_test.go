package main

import (
	"testing"
)

func BenchmarkDupBirth(b *testing.B) {
	var k uint64
	for i := 0; i < b.N; i++ {
		k += uint64(RandDup(366))
	}
	b.Logf("average %d people create duplicate birthday", int(k)/b.N)
}

func BenchmarkCollectAll(b *testing.B) {
	top := 20000
	var k uint64
	for i := 0; i < b.N; i++ {
		k += uint64(CollectAll(top))
	}
	b.Logf("average %d times collect all numbers of %d", int(k)/b.N, top)
}
