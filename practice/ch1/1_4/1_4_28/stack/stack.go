package stack

import "algo/cases/ch1/1_3/queue"

// queue-structured stack
type Stack[T any] struct {
	q queue.Queue[T]
}

type node[T any] struct {
	v    T
	next *node[T]
}

func (s *Stack[T]) Empty() bool {
	return s.q.Size() == 0
}

func (s *Stack[T]) Size() int {
	return s.q.Size()
}

func (s *Stack[T]) Push(v T) {
	s.q.Enqueue(v)
}

func (s *Stack[T]) Pop() *T {
	if s.q.Size() == 0 {
		return nil
	}
	for i := 0; i < s.q.Size()-1; i++ {
		v := s.q.Dequeue()
		s.q.Enqueue(*v)
	}
	return s.q.Dequeue()
}

func (s *Stack[T]) Peek() *T {
	if s.q.Size() == 0 {
		return nil
	}
	var v *T
	for i := 0; i < s.q.Size(); i++ {
		v = s.q.Dequeue()
		s.q.Enqueue(*v)
	}
	return v
}

func (s *Stack[T]) Traverse(f func(T) bool) {
	s.q.Traverse(f)
}
