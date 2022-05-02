package randqueue

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type node[T any] struct {
	v    T
	next *node[T]
}

type RandQueue[T any] struct {
	last  *node[T]
	first *node[T]
	size  int
}

func (q *RandQueue[T]) Size() int {
	return q.size
}

func (q *RandQueue[T]) Enqueue(v T) {
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

// return: chosen node and the node before chosen node
func (q *RandQueue[T]) sample() (*node[T], *node[T]) {
	if q.size == 0 {
		return nil, nil
	}
	pos := rand.Intn(q.size) // [0, q.size)
	if pos == 0 {
		return q.first, nil
	}
	p := q.first
	for i := 0; i <= pos; i++ {
		if i == pos-1 {
			break
		}
		p = p.next
	}
	return p.next, p
}

func (q *RandQueue[T]) Sample() *T {
	p, _ := q.sample()
	if p == nil {
		return nil
	}
	return &p.v
}

func (q *RandQueue[T]) Dequeue() *T {
	p, pre := q.sample()
	if pre == nil && p == nil {
		return nil
	}
	// p will never be nil
	if pre == nil {
		q.first = p.next
	} else {
		if p.next != nil {
			pre.next = p.next
		} else {
			pre.next = nil
		}
	}
	q.size--
	return &p.v
}

// should be provide or not
func (q *RandQueue[T]) Traverse(f func(T) bool) {
	for p := q.first; p != nil; p = p.next {
		if f(p.v) {
			break
		}
	}
}

func (q *RandQueue[T]) DoAndDequeue(f func(T) bool) {
	for {
		v := q.Dequeue()
		if v == nil {
			break
		}
		if f(*v) {
			break
		}
	}
}
