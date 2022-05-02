package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"algo/cases/ch1/1_3/queue"
	"fmt"
	"os"
)

func main() {
	s := linkstack.Stack[string]{}
	q := queue.Queue[string]{}

	for _, v := range os.Args[1:] {
		q.Enqueue(v)
	}

	for q.Size() != 0 {
		s.Push(*q.Dequeue())
	}
	for !s.Empty() {
		q.Enqueue(*s.Pop())
	}
	q.Traverse(func(s string) bool {
		fmt.Printf("%s ", s)
		return false
	})
	fmt.Println()
}
