package queue

type node[T any] struct {
	v    T
	next *node[T]
}

type Queue[T any] struct {
	last *node[T]
}

func (q *Queue[T]) traverse(f func(*node[T]) bool) {
	if q.last == nil {
		return
	}
	for p := q.last.next; p != q.last; p = p.next {
		if f(p) {
			break
		}
	}
	f(q.last)
}

func (q *Queue[T]) Enqueue(v T) {
	if q.last == nil {
		n := &node[T]{
			v:    v,
			next: nil,
		}
		n.next = n
		q.last = n
		return
	}
	n := &node[T]{
		v:    v,
		next: q.last.next,
	}
	q.last.next = n
	q.last = n
}

func (q *Queue[T]) Dequeue() *T {
	if q.last == nil {
		return nil
	}
	if q.last.next == q.last {
		p := q.last
		q.last = nil
		return &p.v
	}
	p := q.last.next
	q.last.next = q.last.next.next
	p.next = nil
	return &p.v
}

func (q *Queue[T]) Size() int {
	cnt := 0
	q.traverse(func(n *node[T]) bool {
		cnt++
		return false
	})
	return cnt
}

func (q *Queue[T]) Traverse(f func(T) bool) {
	q.traverse(func(n *node[T]) bool {
		return f(n.v)
	})
}
