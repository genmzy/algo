package intset

type IntSet []int

func (s IntSet) Contains(n int) bool {
	return s.Rank(n) != -1
}

// practice 1.4.10
func (s IntSet) HowMany(n int) int {
	l := 0
	r := len(s) - 1
	idx := -1
	for l <= r {
		mid := l + (r-l)>>1
		if n < s[mid] {
			r = mid - 1
		} else if n > s[mid] {
			l = mid + 1
		} else {
			idx = mid
			break
		}
	}
	if idx == -1 {
		return 0
	}
	cnt := 0
	for i := idx; i < len(s); i++ {
		if s[i] == n {
			cnt++
		} else {
			break
		}
	}
	for i := idx; i > 0; i-- {
		if s[i] == n {
			cnt++
		} else {
			break
		}
	}
	return cnt - 1
}

// find the index of n in nums by binary search
// if the key is the only one in the set, this will be correct
// if the set have more than one n, will return one of the index
// or return -1
func (s IntSet) Rank(n int) int {
	l := 0
	r := len(s) - 1
	for l <= r {
		mid := l + (r-l)>>1
		if n < s[mid] {
			r = mid - 1
		} else if n > s[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func (s IntSet) MiniIdx(n int) int {
	l := 0
	r := len(s) - 1
	ret := -1
	for l <= r {
		mid := l + (r-l)>>1
		if n < s[mid] {
			r = mid - 1
		} else if n > s[mid] {
			l = mid + 1
		} else {
			ret = mid
			break
		}
	}
	if ret == -1 {
		return ret
	}
	// 2, 4, 5, 5
	for r = ret; l < r; {
		mid := l + (r-l)>>1
		if n > s[mid] {
			l = mid + 1
		} else { // n == nums[mid]
			r = mid
		}
		if l == r {
			ret = r
		}
	}
	return ret
}
