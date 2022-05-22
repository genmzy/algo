package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"fmt"
)

func dump(s *linkstack.Stack[int]) {
	fmt.Print("dumping ")
	s.Traverse(func(i int) bool {
		fmt.Printf("%d ", i)
		return false
	})
	fmt.Println("lenth :", s.Size())
}

func main() {
	s := &linkstack.Stack[int]{}
	for i := 0; i <= 10; i++ {
		s.Push(i)
	}
	dump(s)
	fmt.Println(s.Empty())

	fmt.Printf("popped: ")
	for i := 0; i < 3; i++ {
		if p := s.Pop(); p != nil {
			fmt.Printf("%d ", *p)
		} else {
			break
		}
	}
	fmt.Println()
	dump(s)
	fmt.Println(s.Empty())

	with := &linkstack.Stack[int]{}
	for i := 100; i > 90; i-- {
		with.Push(i)
	}
	dump(with)

	s.Catenation(with)
	dump(s)

	// s link with an empty stack
	s.Catenation(&linkstack.Stack[int]{})
	dump(s)

	// an empty stack link with s
	m := &linkstack.Stack[int]{}
	m.Catenation(s)
	dump(m)
}
