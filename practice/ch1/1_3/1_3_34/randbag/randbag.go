package randbag

func NewRandBag[T comparable]() *RandBag[T] {
	return &RandBag[T]{bag: make(map[T]bool)}
}

type RandBag[T comparable] struct {
	bag map[T]bool
}

func (b *RandBag[T]) Size() int {
	return len(b.bag)
}

func (b *RandBag[T]) Add(v T) {
	b.bag[v] = true
}

func (b *RandBag[T]) Traverse(f func(T) bool) {
	for k := range b.bag {
		if f(k) {
			break
		}
	}
}
