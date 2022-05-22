package main

import (
	"algo/cases/ch1/1_1/binarysearch"
	"sort"
)

func IntPairs(nums []int) int {
	cnt := 0
	sort.Ints(nums)
	l := len(nums)
	for i := 0; i < l; i++ {
		if binarysearch.Rank(nums[i], nums) > i {
			cnt++
		}
	}
	return cnt
}
