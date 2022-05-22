package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"testing"
)

func BenchmarkInt(b *testing.B) {
	s := IntStack{}
	for i := 0; i < b.N; i++ {
		if i%2 == 1 {
			s.Push(i)
			continue
		}
		s.Pop()
	}
}

func BenchmarkGeneric(b *testing.B) {
	s := linkstack.Stack[int]{}
	for i := 0; i < b.N; i++ {
		if i%2 == 1 {
			s.Push(i)
			continue
		}
		s.Pop()
	}
}

func BenchmarkIface(b *testing.B) {
	s := IfaceStack{}
	for i := 0; i < b.N; i++ {
		if i%2 == 1 {
			s.Push(i)
			continue
		}
		s.Pop()
	}
}
