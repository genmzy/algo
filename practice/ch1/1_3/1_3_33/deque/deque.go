package deque

type node[T any] struct {
	v    T
	next *node[T]
	pre  *node[T]
}

type Deque[T any] struct {
	first *node[T]
	last  *node[T]
	size  int
}

func (dq *Deque[T]) Size() int {
	return dq.size
}

func (dq *Deque[T]) PushLeft(v T) {
	n := &node[T]{
		v:    v,
		pre:  nil,
		next: dq.first,
	}
	if dq.size == 0 {
		dq.last = n
	} else {
		dq.first.pre = n
	}
	dq.first = n
	dq.size++
}

func (dq *Deque[T]) PushRight(v T) {
	n := &node[T]{
		v:    v,
		pre:  dq.last,
		next: nil,
	}
	if dq.size == 0 {
		dq.first = n
	} else {
		dq.last.next = n
	}
	dq.last = n
	dq.size++
}

func (dq *Deque[T]) PopLeft() *T {
	if dq.size == 0 {
		return nil
	}
	p := dq.first
	dq.first = p.next
	if dq.size == 1 {
		dq.last = nil
	} else {
		p.next.pre = nil
	}
	dq.size--
	return &p.v
}

func (dq *Deque[T]) PopRight() *T {
	if dq.size == 0 {
		return nil
	}
	p := dq.last
	dq.last = p.pre
	if dq.size == 1 {
		dq.first = nil
	} else {
		p.pre.next = nil
	}
	dq.size--
	return &p.v
}

func (dq *Deque[T]) TraverseFromLeft(f func(T) bool) {
	for p := dq.first; p != nil; p = p.next {
		if f(p.v) {
			break
		}
	}
}

func (dq *Deque[T]) TraverseFromRight(f func(T) bool) {
	for p := dq.last; p != nil; p = p.pre {
		if f(p.v) {
			break
		}
	}
}
