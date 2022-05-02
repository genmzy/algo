package bag

type node[T any] struct {
	v    T
	next *node[T]
}

type Bag[T any] struct {
	top  *node[T]
	size int
}

func (s *Bag[T]) Empty() bool {
	return s.top == nil
}

func (s *Bag[T]) Size() int {
	return s.size
}

func (s *Bag[T]) Add(v T) {
	n := &node[T]{
		v:    v,
		next: s.top,
	}
	s.top = n
	s.size++
}

// traverse from the top to the bottom of the stack
func (s *Bag[T]) Traverse(f func(T) bool) {
	for p := s.top; p != nil; p = p.next {
		if f(p.v) {
			break
		}
	}
}
