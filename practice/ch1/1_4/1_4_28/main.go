package main

import (
	"algo/practice/ch1/1_4/1_4_28/stack"
	"fmt"
)

func dump(s *stack.Stack[int]) {
	fmt.Print("dumping ")
	s.Traverse(func(i int) bool {
		fmt.Printf("%d ", i)
		return false
	})
	fmt.Println()
}

func main() {
	s := &stack.Stack[int]{}
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	dump(s)
	for i := 0; i < 3; i++ {
		p := s.Pop()
		if p != nil {
			fmt.Println(*p)
		} else {
			fmt.Println("nil")
			break
		}
	}
	dump(s)
}
