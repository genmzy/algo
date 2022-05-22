package main

type ifaceNode struct {
	v    interface{}
	next *ifaceNode
}

type IfaceStack struct {
	top  *ifaceNode
	size int
}

func (s *IfaceStack) Empty() bool {
	return s.top == nil
}

func (s *IfaceStack) Size() int {
	return s.size
}

func (s *IfaceStack) Push(v interface{}) {
	n := &ifaceNode{
		v:    v,
		next: s.top,
	}
	s.top = n
	s.size++
}

func (s *IfaceStack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	n := s.top
	s.top = n.next
	n.next = nil // avoid next one dissociated
	s.size--
	return &n.v
}

// 1.3.7
func (s *IfaceStack) Peek() interface{} {
	return s.top.v
}

// traverse from the top to the bottom of the stack
func (s *IfaceStack) Traverse(f func(interface{}) bool) {
	for p := s.top; p != nil; p = p.next {
		if f(p.v) {
			break
		}
	}
}

// 1.3.47 for stack
func (s *IfaceStack) Catenation(with *IfaceStack) {
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
