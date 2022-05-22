package main

type intNode struct {
	v    int
	next *intNode
}

type IntStack struct {
	top  *intNode
	size int
}

func (s *IntStack) Empty() bool {
	return s.top == nil
}

func (s *IntStack) Size() int {
	return s.size
}

func (s *IntStack) Push(v int) {
	n := &intNode{
		v:    v,
		next: s.top,
	}
	s.top = n
	s.size++
}

const intMin = ^int(^uint(0) >> 1) // `-2^n` for n bit machine

func (s *IntStack) Pop() int {
	if s.size == 0 {
		return intMin
	}
	n := s.top
	s.top = n.next
	n.next = nil // avoid next one dissociated
	s.size--
	return n.v
}

// 1.3.7
func (s *IntStack) Peek() *int {
	return &s.top.v
}

// traverse from the top to the bottom of the stack
func (s *IntStack) Traverse(f func(int) bool) {
	for p := s.top; p != nil; p = p.next {
		if f(p.v) {
			break
		}
	}
}

// 1.3.47 for stack
func (s *IntStack) Catenation(with *IntStack) {
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
