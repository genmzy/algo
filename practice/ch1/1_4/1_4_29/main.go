package main

import (
	"algo/practice/ch1/1_4/1_4_29/steque"
	"fmt"
)

func main() {
	sq := &steque.Steque[int]{}

	for i := 0; i < 10; i++ {
		if i == 5 {
			sq.Enqueue(i)
		}
		sq.Push(i)
	}
	//dump(sq)
	for i := 0; i < 5; i++ {
		fmt.Println("first popping ", *sq.Pop())
	}
	//dump(sq)
	for i := 0; i < 10; i++ {
		sq.Enqueue(i)
	}
	//dump(sq)
	size := sq.Size()
	for i := 0; i < size; i++ {
		fmt.Println("second popping ", *sq.Pop())
	}
}
