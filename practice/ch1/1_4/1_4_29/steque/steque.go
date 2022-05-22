package steque

import "algo/cases/ch1/1_3/linkstack"

// pop   n(k)
//        ↑
//       n(k) -> n(k-1) -> n(k-2) -> ... -> n(2) -> n(1)
//        ↑                                          ↓
// push n(k+1)                             enqueue: n(0)

type Steque[T any] struct {
	m linkstack.Stack[T]
	n linkstack.Stack[T]
}

func (sq *Steque[T]) reverseM() {
	for sq.m.Size() != 0 {
		sq.n.Push(*sq.m.Pop())
	}
}

func (sq *Steque[T]) reverseN() {
	for sq.n.Size() != 0 {
		sq.m.Push(*sq.n.Pop())
	}
}

func (sq *Steque[T]) Push(v T) {
	sq.reverseN()
	sq.m.Push(v)
}

func (sq *Steque[T]) Pop() *T {
	if sq.Size() == 0 {
		return nil
	}
	sq.reverseN()
	return sq.m.Pop()
}

func (sq *Steque[T]) Enqueue(v T) {
	sq.reverseM()
	sq.n.Push(v)
}

func (sq *Steque[T]) Size() int {
	return sq.m.Size() + sq.n.Size()
}
