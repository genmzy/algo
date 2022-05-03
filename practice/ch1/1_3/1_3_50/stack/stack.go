package stack

import "errors"

var ErrConcurrentModified = errors.New("concurrent modifications")

type node[T any] struct {
	v    T
	next *node[T]
}

type Stack[T any] struct {
	top    *node[T]
	size   int
	popped int
	pushed int
}

func (s *Stack[T]) Empty() bool {
	return s.top == nil
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) Push(v T) {
	n := &node[T]{
		v:    v,
		next: s.top,
	}
	s.top = n
	s.size++
	s.pushed++
}

func (s *Stack[T]) Pop() *T {
	if s.size == 0 {
		return nil
	}
	n := s.top
	s.top = n.next
	n.next = nil // avoid next one dissociated
	s.size--
	s.popped++
	return &n.v
}

// 1.3.7
func (s *Stack[T]) Peek() *T {
	return &s.top.v
}

// traverse from the top to the bottom of the stack
func (s *Stack[T]) Traverse(f func(T) bool) {
	popped := s.popped
	pushed := s.pushed
	for p := s.top; p != nil; p = p.next {
		if s.popped+s.size != s.pushed || s.popped != popped || s.pushed != pushed {
			panic(ErrConcurrentModified)
		}
		if f(p.v) {
			break
		}
	}
}

// 1.3.47 for stack
func (s *Stack[T]) Catenation(with *Stack[T]) {
	if with.size == 0 {
		return
	}
	// find the bottom of stack with
	p := with.top
	for {
		if p.next == nil {
			break
		}
		p = p.next
	}
	p.next = s.top
	s.size += with.size
	s.top = with.top
}
