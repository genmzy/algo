package binarysearch

// find the index of n in nums by binary search
// if the key is the only one in the set, this will be correct
// if the set have more than one n, will return one of the index
// or return -1
func Rank(n int, nums []int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)>>1
		if n < nums[mid] {
			r = mid - 1
		} else if n > nums[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func RankDesc(n int, nums []int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)>>1
		if n < nums[mid] {
			l = mid + 1
		} else if n > nums[mid] {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// practice 1.4.10
func FindMinIdx(n int, nums []int) int {
	l := 0
	r := len(nums) - 1
	ret := -1
	for l <= r {
		mid := l + (r-l)>>1
		if n < nums[mid] {
			r = mid - 1
		} else if n > nums[mid] {
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
		if n > nums[mid] {
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
