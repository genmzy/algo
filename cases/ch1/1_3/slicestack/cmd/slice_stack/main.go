package main

import (
	"algo/cases/ch1/1_3/slicestack"
	"fmt"
)

func main() {
	s := slicestack.Stack[int]{}
	for i := 0; i < 10; i++ {
		s.Push(i)
		s.Traverse(func(v int) bool {
			fmt.Printf("%d ", v)
			return false
		})
		fmt.Println()
	}
}
