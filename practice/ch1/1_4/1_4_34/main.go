package main

import "fmt"

type Comp int

const (
	Eq Comp = iota
	Gt
	Lt
)

func GuessNumber(down, up int, f func(int) Comp) int {
	if down > up {
		return -1
	}
	for down <= up {
		mid := down + (up-down)>>1
		switch f(mid) {
		case Gt:
			up = mid - 1
		case Lt:
			down = mid + 1
		case Eq:
			return mid
		default:
			panic("invalid result")
		}
	}
	return -1
}

func main() {
	fmt.Println(GuessNumber(0, 100, func(i int) Comp {
		if i > 20 {
			return Gt
		}
		if i < 20 {
			return Lt
		}
		return Eq
	}))
}
