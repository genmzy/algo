package queue

type node[T any] struct {
	v    T
	next *node[T]
}

type Queue[T any] struct {
	last  *node[T]
	first *node[T]
	size  int
}

func (q *Queue[T]) Enqueue(v T) {
	n := &node[T]{
		v:    v,
		next: nil,
	}
	if q.last != nil {
		q.last.next = n
	}
	if q.size == 0 {
		q.first = n
	}
	q.last = n
	q.size++
}

func (q *Queue[T]) Dequeue() *T {
	if q.size == 0 {
		return nil
	}
	n := q.first
	q.first = n.next
	if q.first == nil {
		q.last = nil
	}
	n.next = nil // avoid next one dissociated
	q.size--
	return &n.v
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Traverse(f func(T) bool) {
	for p := q.first; p != nil; p = p.next {
		if f(p.v) {
			break
		}
	}
}
