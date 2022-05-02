package main

import (
	"algo/practice/ch1/1_3/1_3_32/steque"
	"fmt"
)

func dump(sq *steque.Steque[int]) {
	fmt.Print("dumping ")
	sq.Traverse(func(v int) bool {
		fmt.Printf("%d ", v)
		return false
	})
	fmt.Println()
}

func main() {
	sq := &steque.Steque[int]{}

	for i := 0; i < 10; i++ {
		sq.Push(i)
	}
	dump(sq)
	size := sq.Size()
	for i := 0; i < size; i++ {
		fmt.Println("popping ", *sq.Pop())
	}
	dump(sq)
	for i := 0; i < 10; i++ {
		sq.Enqueue(i)
	}
	dump(sq)
	size = sq.Size()
	for i := 0; i < size; i++ {
		fmt.Println("popping ", *sq.Pop())
	}
	dump(sq)
}
