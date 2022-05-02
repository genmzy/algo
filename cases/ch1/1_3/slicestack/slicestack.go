package slicestack

type Stack[T any] struct {
	base []T
}

func (s *Stack[T]) Push(v T) {
	s.base = append(s.base, v)
}

func (s *Stack[T]) Pop() *T {
	if len(s.base) == 0 {
		return nil
	}
	// NOTE: this will add GC pressure
	// res := s.base[len(s.base)-1]
	// s.base = s.base[:len(s.base)-1]
	// return &res
	res := &s.base[len(s.base)-1]
	s.base = s.base[:len(s.base)-1]
	return res
}

func (s *Stack[T]) Size() int {
	return len(s.base)
}

func (s *Stack[T]) Top() *T {
	return &s.base[len(s.base)-1]
}

func (s *Stack[T]) Peek() *T {
	return &s.base[len(s.base)-1]
}

// traverse all elements by call function `f`, if f return true, that
// means break
func (s *Stack[T]) Traverse(f func(T) bool) {
	for _, x := range s.base {
		if f(x) {
			break
		}
	}
}
