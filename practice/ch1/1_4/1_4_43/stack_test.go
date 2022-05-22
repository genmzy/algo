package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"algo/cases/ch1/1_3/slicestack"
	"testing"
)

var k int = 20000000

func BenchmarkLink(b *testing.B) {
	s := linkstack.Stack[int]{}
	for i := 0; i < b.N; i++ {
		if i > k {
			break
		}
		s.Push(i)
	}
}

func BenchmarkStack(b *testing.B) {
	s := slicestack.Stack[int]{}
	for i := 0; i < b.N; i++ {
		if i > k {
			break
		}
		s.Push(i)
	}
}
