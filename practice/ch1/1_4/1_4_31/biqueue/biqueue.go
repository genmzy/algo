package biqueue

import "algo/cases/ch1/1_3/linkstack"

// TODO: finish later <2022-05-21, genmzy> //
type Biqueue[T any] struct {
	a, b, c linkstack.Stack[T]
}

func (q *Biqueue[T]) reverseA() {

}

func (q *Biqueue[T]) reverseB() {

}

func (q *Biqueue[T]) reverseC() {

}

func (q *Biqueue[T]) LeftEnqueue(v T) {
}

func (q *Biqueue[T]) LeftDequeue() *T {
	return nil
}

func (q *Biqueue[T]) RightEnqueue(v T) {
}

func (q *Biqueue[T]) RightDequeue() *T {
	return nil
}

func (q *Biqueue[T]) Size() int {
	return q.a.Size() + q.b.Size() + q.c.Size()
}
