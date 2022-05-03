package steque

type node[T any] struct {
	v    T
	next *node[T]
}

type Steque[T any] struct {
	top  *node[T]
	tail *node[T]
	size int
}

func (sq *Steque[T]) Push(v T) {
	n := &node[T]{
		v:    v,
		next: sq.top,
	}
	sq.top = n
	if sq.size == 0 {
		sq.tail = n
	}
	sq.size++
}

func (sq *Steque[T]) Pop() *T {
	if sq.size == 0 {
		return nil
	}
	p := sq.top
	sq.top = p.next
	if sq.size == 1 {
		sq.tail = nil
	}
	sq.size--
	p.next = nil
	return &p.v
}

func (sq *Steque[T]) Enqueue(v T) {
	n := &node[T]{
		v:    v,
		next: nil,
	}
	if sq.size != 0 {
		sq.tail.next = n
	} else {
		sq.top = n
	}
	sq.tail = n
	sq.size++
}

func (sq *Steque[T]) Traverse(f func(v T) bool) {
	for n := sq.top; n != nil; n = n.next {
		if f(n.v) {
			break
		}
	}
}

func (sq *Steque[T]) Size() int {
	return sq.size
}

func (sq *Steque[T]) Catenation(with *Steque[T]) {
	if with.size == 0 {
		return
	}
	if sq.size == 0 {
		sq.top = with.top
	} else {
		sq.tail.next = with.top
	}
	sq.tail = with.tail
	sq.size += with.size
}
