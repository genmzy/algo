package main

import (
	"algo/practice/ch1/1_4/1_4_27/queue"
	"fmt"
)

func dump(q *queue.Queue[int]) {
	fmt.Printf("dumping result: ")
	q.Traverse(func(i int) bool {
		fmt.Printf("%d ", i)
		return false
	})
	fmt.Println()
}

func main() {
	q := &queue.Queue[int]{}
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	dump(q)
	for i := 0; i < 3; i++ {
		p := q.Dequeue()
		if p != nil {
			fmt.Println("popped:", *p)
		} else {
			fmt.Println("nil")
		}
	}
	dump(q)
}
