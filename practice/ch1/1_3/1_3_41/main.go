package main

import (
	"algo/cases/ch1/1_3/queue"
	"fmt"
	"strconv"
)

func CopyQueue[T any](q *queue.Queue[T]) *queue.Queue[T] {
	ret := &queue.Queue[T]{}
	q.Traverse(func(v T) bool {
		ret.Enqueue(v)
		return false
	})
	return q
}

type Stringer interface {
	String() string
}

func dump[T Stringer](q *queue.Queue[T]) {
	fmt.Println("dumping ")
	q.Traverse(func(v T) bool {
		fmt.Print(v.String(), " ")
		return false
	})
	fmt.Println()
}

type mytype int

func (t mytype) String() string {
	return strconv.Itoa(int(t))
}

func main() {
	q := &queue.Queue[mytype]{}
	for i := mytype(0); i < 10; i++ {
		q.Enqueue(i)
	}
	dump(q)

	ret := CopyQueue(q)
	dump(ret)
}
