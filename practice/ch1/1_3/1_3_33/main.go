package main

import (
	"algo/practice/ch1/1_3/1_3_33/deque"
	"fmt"
)

func dump(dq *deque.Deque[int]) {
	fmt.Print("dumping ")
	dq.TraverseFromLeft(func(i int) bool {
		fmt.Printf("%d ", i)
		return false
	})
	fmt.Println()
}

func main() {
	dq := &deque.Deque[int]{}
	for i := 0; i < 10; i++ {
		dq.PushLeft(i)
	}
	dump(dq)
	size := dq.Size()
	for i := 0; i < size; i++ {
		fmt.Println("popping:", *dq.PopLeft())
	}
	dump(dq)

	for i := 0; i < 10; i++ {
		dq.PushRight(i)
	}
	dump(dq)
	size = dq.Size()
	for i := 0; i < size; i++ {
		fmt.Println("popping:", *dq.PopRight())
	}
	dump(dq)
}
