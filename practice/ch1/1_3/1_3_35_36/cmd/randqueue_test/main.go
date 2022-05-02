package main

import (
	"algo/practice/ch1/1_3/1_3_35_36/randqueue"
	"fmt"
)

func dump(q *randqueue.RandQueue[int]) {
	fmt.Print("dumping ")
	q.Traverse(func(i int) bool {
		fmt.Printf("%d ", i)
		return false
	})
	fmt.Println()
}

func main() {
	q := &randqueue.RandQueue[int]{}
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	dump(q)

	fmt.Println("simple generate random index:")
	for i := 0; i < 50; i++ {
		fmt.Printf("%d ", *q.Sample())
	}
	fmt.Println()

	fmt.Println("dequeue all elements")
	for i := 0; i < 50; i++ {
		p := q.Dequeue()
		if p == nil {
			fmt.Print("<nil> ")
			continue
		}
		fmt.Printf("%d ", *p)
	}
	fmt.Println()
}
