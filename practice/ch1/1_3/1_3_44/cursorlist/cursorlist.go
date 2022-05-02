package cursorlist

type binode[T any] struct {
	v    T
	pre  *binode[T]
	next *binode[T]
}

// a general link list with cursor
type CursorList[T any] struct {
	first  *binode[T]
	cursor *binode[T]
	last   *binode[T]
	size   int
}

func (l *CursorList[T]) InsertBefore(v T) {
	defer func() { l.size++ }()
	n := &binode[T]{
		v:    v,
		next: l.cursor,
		pre:  nil,
	}
	if l.size == 0 {
		l.first = n
		l.cursor = n
		l.last = n
		return
	}
	if l.cursor.pre != nil {
		l.cursor.pre.next = n
		n.pre = l.cursor.pre
	} else {
		l.first = n
	}
	l.cursor.pre = n
}

func (l *CursorList[T]) InsertAfter(v T) {
	defer func() { l.size++ }()
	n := &binode[T]{
		v:    v,
		pre:  l.cursor,
		next: nil,
	}
	if l.size == 0 {
		l.first = n
		l.cursor = n
		return
	}
	if l.cursor.next != nil {
		l.cursor.next.pre = n
		n.next = l.cursor.next
	} else {
		l.last = n
	}
	l.cursor.next = n
}

func (l *CursorList[T]) InsertBeforeN(v T, n int) {
	if n < 0 {
		return
	}
	cursor := l.cursor
	for i := 0; i < n-1; i++ {
		if l.cursor == nil {
			l.cursor = cursor
			return
		}
		l.cursor = l.cursor.pre
	}
	l.InsertBefore(v)
	if l.size != 0 {
		l.cursor = cursor
	}
}

func (l *CursorList[T]) InsertAfterN(v T, n int) {
	if n < 0 {
		return
	}
	cursor := l.cursor
	for i := 0; i < n-1; i++ {
		if l.cursor == nil {
			l.cursor = cursor
			return
		}
		l.cursor = l.cursor.pre
	}
	l.InsertBefore(v)
	if l.size != 0 {
		l.cursor = cursor
	}
}

func (l *CursorList[T]) DeleteAfterN(n int) *T {
	if n < 0 || l.size == 0 {
		return nil
	}
	cursor := l.cursor
	// find the delete one
	for i := 0; i < n; i++ {
		if l.cursor == nil {
			l.cursor = cursor
			return nil
		}
		l.cursor = l.cursor.next
	}
	ret := l.Delete()
	if n != 0 {
		l.cursor = cursor
	}
	return ret
}

func (l *CursorList[T]) DeleteBeforeN(n int) *T {
	if n < 0 || l.size == 0 {
		return nil
	}
	cursor := l.cursor
	// find the one that should be delete
	for i := 0; i < n; i++ {
		if l.cursor == nil {
			l.cursor = cursor
			return nil
		}
		l.cursor = l.cursor.pre
	}
	ret := l.Delete()
	if n != 0 {
		l.cursor = cursor
	}
	return ret
}

func (l *CursorList[T]) Delete() *T {
	if l.size == 0 {
		return nil
	}
	p := l.cursor
	if p.pre != nil {
		p.pre.next = p.next
	} else {
		l.first = p.next
	}
	if p.next != nil {
		p.next.pre = p.pre
	} else {
		l.last = p.pre
	}
	l.size--
	return &p.v
}

func (l *CursorList[T]) left() {
	if l.size == 0 || l.cursor.pre == nil {
		return
	}
	l.cursor = l.cursor.pre
}

func (l *CursorList[T]) right() {
	if l.size == 0 || l.cursor.next == nil {
		return
	}
	l.cursor = l.cursor.next
}

func (l *CursorList[T]) Left(k int) {
	for i := 0; i < k; i++ {
		l.left()
	}
}

func (l *CursorList[T]) Right(k int) {
	for i := 0; i < k; i++ {
		l.right()
	}
}

func (l *CursorList[T]) Size() int {
	return l.size
}

func (l *CursorList[T]) Cursor() *T {
	if l.size == 0 {
		return nil
	}
	return &l.cursor.v
}

func (l *CursorList[T]) CursorBeforeN(n int) *T {
	if l.size == 0 {
		return nil
	}
	cursor := l.cursor
	defer func() {
		l.cursor = cursor
	}()
	for i := 0; i < n; i++ {
		if l.cursor == nil {
			return nil
		}
		l.cursor = l.cursor.pre
	}
	ret := &l.cursor.v
	return ret
}

func (l *CursorList[T]) CursorAfterN(n int) *T {
	if l.size == 0 {
		return nil
	}
	cursor := l.cursor
	defer func() {
		l.cursor = cursor
	}()
	for i := 0; i < n; i++ {
		if l.cursor == nil {
			return nil
		}
		l.cursor = l.cursor.next
	}
	ret := &l.cursor.v
	return ret
}

func (l *CursorList[T]) Add(v T) {
	cursor := l.cursor
	for cursor != nil {
		if l.cursor.next == nil {
			break
		}
		l.cursor = l.cursor.next
	}
	l.InsertAfter(v)
	if cursor != nil {
		l.cursor = cursor
	}
}

func (l *CursorList[T]) traverse(f func(n *binode[T]) bool, from *binode[T], desc bool) {
	if from != l.cursor && from != l.first && from != l.last {
		return
	}
	for p := from; p != nil; {
		if f(p) {
			break
		}
		if desc {
			p = p.pre
		} else {
			p = p.next
		}
	}
}

// can change original value of linklist, see practise 1_3_44
func (l *CursorList[T]) AllDo(f func(v *T) bool) {
	l.traverse(func(n *binode[T]) bool {
		return f(&n.v)
	}, l.first, false)
}

func (l *CursorList[T]) Traverse(f func(v T) bool) {
	l.traverse(func(n *binode[T]) bool {
		return f(n.v)
	}, l.first, false)
}

func (l *CursorList[T]) TraverseDesc(f func(v T) bool) {
	l.traverse(func(n *binode[T]) bool {
		return f(n.v)
	}, l.last, true)
}

func (l *CursorList[T]) TraverseFromCursor(f func(v T) bool, desc bool) {
	l.traverse(func(n *binode[T]) bool {
		return f(n.v)
	}, l.cursor, desc)
}
