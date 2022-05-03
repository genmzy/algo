package dstack

import "algo/practice/ch1/1_3/1_3_33/deque"

type DStack[T any] struct {
	deque.Deque[T]
}
