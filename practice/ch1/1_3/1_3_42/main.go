package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"fmt"
	"strconv"
)

func CopyQueue[T any](q *linkstack.Stack[T]) *linkstack.Stack[T] {
	ret := &linkstack.Stack[T]{}
	q.Traverse(func(v T) bool {
		ret.Push(v)
		return false
	})
	return q
}

type Stringer interface {
	String() string
}

func dump[T Stringer](q *linkstack.Stack[T]) {
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
	q := &linkstack.Stack[mytype]{}
	for i := mytype(0); i < 10; i++ {
		q.Push(i)
	}
	dump(q)

	ret := CopyQueue(q)
	dump(ret)
}
