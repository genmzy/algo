package main

import (
	"algo/practice/ch1/1_3/1_3_29_39/queue"
	"fmt"
)

func dump(q queue.Queue[int]) {
	q.Traverse(func(i int) bool {
		fmt.Print(i, " ")
		return false
	})
	fmt.Println()
}

func main() {
	q := queue.Queue[int]{}
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	dump(q)
	sz := q.Size()
	fmt.Println("size:", sz)
	//for i := 0; i < q.Size(); i++ { // q.Size() always change because q.Dequeue
	for i := 0; i < sz; i++ {
		v := q.Dequeue()
		if v != nil {
			fmt.Println(*v)
		} else {
			fmt.Println("got nil")
		}
	}
}
