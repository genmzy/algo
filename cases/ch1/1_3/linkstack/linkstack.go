package linkstack

type node[T any] struct {
	v    T
	next *node[T]
}

type Stack[T any] struct {
	top  *node[T]
	size int
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
}

func (s *Stack[T]) Pop() *T {
	if s.size == 0 {
		return nil
	}
	n := s.top
	s.top = n.next
	n.next = nil // avoid next one dissociated
	s.size--
	return &n.v
}

func (s *Stack[T]) Top() *T {
	return &s.top.v
}

// traverse from the top to the bottom of the stack
func (s *Stack[T]) Traverse(f func(T) bool) {
	for p := s.top; p != nil; p = p.next {
		if f(p.v) {
			break
		}
	}
}
