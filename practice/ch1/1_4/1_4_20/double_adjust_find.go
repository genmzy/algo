package main

import (
	"algo/cases/ch1/1_1/binarysearch"
	"fmt"
)

func DoubleAdjustableFind(nums []int, n int) []int {
	// find the critial point
	res := make([]int, 0)
	edge := false
	if n == nums[len(nums)-1] {
		edge = true
		res = append(res, len(nums)-1)
	}
	if n == nums[0] {
		edge = true
		res = append(res, 0)
	}
	if edge {
		return res
	}
	l, r := 1, len(nums)-2
	critial := -1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid+1] > nums[mid] && nums[mid] > nums[mid-1] {
			l = mid + 1
		} else if nums[mid+1] < nums[mid] && nums[mid] < nums[mid-1] {
			r = mid - 1
		} else {
			critial = mid
			break
		}
	}
	if critial == -1 {
		return res
	}
	i := binarysearch.Rank(n, nums[:critial])
	if i != -1 {
		res = append(res, i)
	}
	i = binarysearch.RankDesc(n, nums[critial:])
	if i != -1 {
		res = append(res, critial+i)
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 4, 3, 1}
	fmt.Println(DoubleAdjustableFind(nums, 3))
}
