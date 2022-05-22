package queue

import (
	"algo/cases/ch1/1_3/linkstack"
)

type node[T any] struct {
	v    T
	next *node[T]
}

// two-link-stack-structured-queue
type Queue[T any] struct {
	s, b linkstack.Stack[T]
}

func (q *Queue[T]) Enqueue(v T) {
	q.b.Push(v)
}

func (q *Queue[T]) Dequeue() *T {
	if q.Size() == 0 {
		return nil
	}
	if q.s.Empty() {
		q.update()
	}
	p := q.s.Pop()
	return p
}

func (q *Queue[T]) update() {
	for !q.b.Empty() {
		p := q.b.Pop()
		q.s.Push(*p)
	}
}

func (q *Queue[T]) Size() int {
	return q.b.Size() + q.s.Size()
}

func (q *Queue[T]) Traverse(f func(T) bool) {
	q.update()
	q.s.Traverse(f)
}

func (q *Queue[T]) Catenation(with *Queue[T]) {
	if with.Size() == 0 {
		return
	}
	q.update()
	with.update()
	q.b = with.s
}
