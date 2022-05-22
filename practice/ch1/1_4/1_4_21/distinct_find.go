package distinctfind

import (
	"algo/cases/ch1/1_1/binarysearch"
	"sort"
)

func DistinctFind(nums []int, n int) int {
	sort.Ints(nums)
	return binarysearch.Rank(n, nums)
}
