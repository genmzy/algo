package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"algo/practice/ch1/1_3/1_3_32/steque"
)

// stack with steque structured bidirect-queue
type Biqueue[T any] struct {
	s  linkstack.Stack[T]
	sq steque.Steque[T]
}

func (q *Biqueue[T]) update() {
	for q.s.Size() != 0 {
		q.sq.Push(*q.s.Pop())
	}
}

func (q *Biqueue[T]) LeftEnqueue(v T) {
	q.s.Push(v)
}

func (q *Biqueue[T]) LeftDequeue() *T {
	if q.Size() == 0 {
		return nil
	}
	if q.sq.Size() == 0 {
		q.update()
	}
	return q.sq.Pop()
}

func (q *Biqueue[T]) RightEnqueue(v T) {
	if q.s.Empty() {
		q.sq.Enqueue(v)
	} else {
		q.s.Push(v)
	}
}

func (q *Biqueue[T]) RightDequeue() *T {
	if q.Size() == 0 {
		return nil
	}
	if q.sq.Size() == 0 {
		q.update()
	}
	return q.sq.Pop()
}

func (q *Biqueue[T]) Size() int {
	return q.s.Size() + q.sq.Size()
}
