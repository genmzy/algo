package fibsearch

var store = make(map[int]int)

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if m, ok := store[n]; ok {
		return m
	}
	m := Fibonacci(n-1) + Fibonacci(n-2)
	store[n] = m
	return m
}

func fibVal(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a = a + b
		b = a - b
	}
	return a
}

// nums is sorted
func FibonacciRank(nums []int, n int) int {
	l, fi := 0, 0
	r := len(nums) - 1
	for fibVal(fi)*2 < r {
		fi++
	}
	for r >= l {
		fi--
		mid := l + fibVal(fi)
		if mid > r {
			mid = r
		}
		if nums[mid] > n {
			r = mid - 1
		} else if nums[mid] < n {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func FibRank(nums []int, k int) int {
	n := len(nums)
	cur, pre, prePre := 1, 1, 0
	for cur <= n {
		cur = cur + pre
		pre = pre + prePre
		prePre = cur - pre
	}
	l := 0
	mid := n - 1
	for prePre >= 0 {
		mid = l + prePre
		if nums[mid] < k {
			l = l + prePre
		} else if nums[mid] == k {
			return mid
		}
		cur = pre
		pre = prePre
		prePre = cur - pre
	}
	return -1
}
