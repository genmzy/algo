package linklist

type node[T comparable] struct {
	v    T
	next *node[T]
}

// a no-general link list
type LinkList[T comparable] struct {
	first *node[T]
}

func (l *LinkList[T]) traverse(f func(*node[T]) bool) {
	for p := l.first; p != nil; p = p.next {
		if f(p) {
			break
		}
	}
}

func (l *LinkList[T]) HeadAdd(v T) {
	n := &node[T]{
		v:    v,
		next: l.first,
	}
	l.first = n
}

func (l *LinkList[T]) Add(v T) {
	n := &node[T]{
		v:    v,
		next: nil,
	}
	if l.first == nil {
		l.first = n
		return
	}
	l.traverse(func(p *node[T]) bool {
		if p.next == nil {
			p.next = n
			return true
		}
		return false
	})
}

func (l *LinkList[T]) DeleteTail() {
	l.traverse(func(p *node[T]) bool {
		if p.next == nil {
			l.first = nil
			return true
		}
		if p.next.next == nil {
			p.next = nil
			return true
		}
		return false
	})
}

func (l *LinkList[T]) Delete(k int) {
	i := 1
	l.traverse(func(p *node[T]) bool {
		if i == k-1 && p.next != nil {
			n := p.next
			p.next = n.next
			n.next = nil
			return true
		}
		i++
		return false
	})
}

func (l *LinkList[T]) findNode(v T) *node[T] {
	var p *node[T]
	l.traverse(func(n *node[T]) bool {
		if n.v == v {
			p = n
			return true
		}
		return false
	})
	return p
}

func (l *LinkList[T]) Find(v T) bool {
	if l.findNode(v) == nil {
		return false
	}
	return true
}

func (l *LinkList[T]) RemoveAfter(v T) {
	n := l.findNode(v)
	if n == nil || n.next == nil {
		return
	}
	n.next = n.next.next
}

func (l *LinkList[T]) InsertAfter(v1, v2 T) {
	n := l.findNode(v1)
	if n == nil {
		return
	}
	newN := &node[T]{
		v:    v2,
		next: n.next,
	}
	n.next = newN
}

func (l *LinkList[T]) Traverse(f func(T) bool) {
	l.traverse(func(n *node[T]) bool {
		return f(n.v)
	})
}

// remove all v(s) in list
func (l *LinkList[T]) Remove(v T) {
	l.traverse(func(n *node[T]) bool {
		if n.next == nil { // the last one
			return true
		}
		if n.next.v == v {
			var last *node[T]
			// avoid continuous number, avoid skip the second like `1, 2, 3, 8, 8, 9` if just use `n.next = n.next.next`
			for p := n.next; p != nil; p = p.next {
				if p.v != v {
					break
				}
				last = p
			}
			n.next = last.next
		}
		return false
	})
	// NOTE: 1. the things we did above ignore the l.first
	//       2. why AFTER the for loop? To avoid meeting: want to delete 8, but got `8, 8, 8, 8, 1, 2, 3, 8, 8`
	if l.first.v == v {
		l.first = l.first.next
	}
}

const intMin = ^int(^uint(0) >> 1) // `-2^n` for n bit machine

func Max(l *LinkList[int]) int {
	i := intMin
	l.traverse(func(n *node[int]) bool {
		if n.v > i {
			i = n.v
		}
		return false
	})
	return i
}

func max(n *node[int], hold *int) {
	if *hold < n.v {
		*hold = n.v
	}
	if n.next != nil {
		max(n.next, hold)
	}
	return
}

func MaxByRecursive(l *LinkList[int]) int {
	i := intMin
	p := l.first
	max(p, &i)
	return i
}

// 1.3.30
func reverse[T comparable](first *node[T]) *node[T] {
	var reverse *node[T]
	for first != nil {
		second := first.next
		first.next = reverse
		reverse = first
		first = second
	}
	return reverse
}

func (l *LinkList[T]) Reverse() {
	l.first = reverse(l.first)
}

// 1.3.30
func revByRec[T comparable](first *node[T]) *node[T] {
	if first == nil {
		return nil
	}
	if first.next == nil {
		return first
	}
	second := first.next
	rest := revByRec(second)
	second.next = first
	first.next = nil
	return rest
}

func (l *LinkList[T]) ReverseByRecursive() {
	l.first = revByRec(l.first)
}
