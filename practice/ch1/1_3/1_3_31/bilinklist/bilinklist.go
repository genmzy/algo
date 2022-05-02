package bilinklist

type node[T comparable] struct {
	v    T
	pre  *node[T]
	next *node[T]
}

type BiLinkList[T comparable] struct {
	head *node[T]
	tail *node[T]
	size int
}

func (l *BiLinkList[T]) HeadAdd(v T) {
	n := &node[T]{
		v:    v,
		pre:  nil,
		next: l.head,
	}
	if l.size == 0 { // head and tail both nil
		l.tail = n
	} else {
		l.head.pre = n
	}
	l.head = n
	l.size++
}

func (l *BiLinkList[T]) TailAdd(v T) {
	n := &node[T]{
		v:    v,
		pre:  l.tail,
		next: nil,
	}
	if l.size == 0 {
		l.head = n
	} else {
		l.tail.next = n
	}
	l.tail = n
	l.size++
}

func (l *BiLinkList[T]) HeadDelete() *T {
	if l.size == 0 {
		return nil
	}
	p := l.head
	l.head = p.next
	if l.size == 1 {
		l.tail = nil
	}
	l.size--
	return &p.v
}

func (l *BiLinkList[T]) TailDelete() *T {
	if l.size == 0 {
		return nil
	}
	p := l.tail
	l.tail = p.pre
	if l.size == 1 {
		l.head = nil
	}
	l.size--
	return &p.v
}

func (l *BiLinkList[T]) traverse(f func(n *node[T]) bool, desc bool) {
	if !desc {
		for p := l.head; p != nil; p = p.next {
			if f(p) {
				break
			}
		}
	} else {
		for p := l.tail; p != nil; p = p.pre {
			if f(p) {
				break
			}
		}
	}
}

func (l *BiLinkList[T]) AddBefore(v T, match T) {
	l.traverse(func(n *node[T]) bool {
		if n.v == match {
			pre := n.pre
			p := &node[T]{
				v:    v,
				pre:  pre,
				next: n,
			}
			if pre != nil {
				pre.next = p
			} else {
				l.head = p
			}
			n.pre = p
			l.size++
			return true
		}
		return false
	}, false)
}

func (l *BiLinkList[T]) AddAfter(v T, match T) {
	l.traverse(func(n *node[T]) bool {
		if n.v == match {
			next := n.next
			p := &node[T]{
				v:    v,
				pre:  n,
				next: next,
			}
			if next != nil {
				next.pre = p
			} else {
				l.tail = p
			}
			n.next = p
			l.size++
			return true
		}
		return false
	}, false)
}

func (l *BiLinkList[T]) Delete(match T) *T {
	var ret *node[T]
	l.traverse(func(n *node[T]) bool {
		if n.v == match {
			pre := n.pre
			next := n.next
			if next != nil {
				next.pre = pre
			} else {
				l.tail = pre
			}
			if pre != nil {
				pre.next = next
			} else {
				l.head = next
			}
			l.size--
			ret = n
			return true
		}
		return false
	}, false)
	return &ret.v
}

func (l *BiLinkList[T]) Traverse(f func(v T) bool, desc bool) {
	l.traverse(func(n *node[T]) bool {
		return f(n.v)
	}, desc)
}
