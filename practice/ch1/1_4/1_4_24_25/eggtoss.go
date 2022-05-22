package eggtoss

import "math"

// f to judge if the egg can be broken or not in k floor
func EggsToss(n int, broke func(k int) bool) int {
	l, r := 0, n
	for l <= r {
		mid := l + (r-l)>>1
		if broke(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if l > n {
		return -1 // if `F' is bigger than `n'
	} else {
		return l
	}
}

func EggsTossAnother(n int, broke func(k int) bool) int {
	l, r := 0, 1
	for !broke(r) && r <= n {
		l = r
		r *= 2
	}
	if r > n {
		return -1
	}
	for l <= r {
		mid := l + (r-l)>>1
		if broke(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

// TODO: figure this out <17-05-22, genmzy> //
func TwoEggsToss(n int, broke func(k int) bool) int {
	step := int(math.Sqrt(float64(n)))
	floor := step
	for !broke(floor) && floor <= n {
		floor += step
	}
	floor -= step
	for !broke(floor) && floor <= n {
		floor++
	}
	return floor
}

// TODO: figure this out <17-05-22, genmzy> //
func TwoEggsTossAnother(n int, broke func(k int) bool) int {
	floor := 2
	for !broke(floor) {
		floor *= floor
	}
	floor = int(math.Sqrt(float64(floor))) - 1
	for !broke(floor) {
		floor++
	}
	return floor
}
