package threesum

import (
	"algo/cases/ch1/1_1/binarysearch"
	"sort"
)

func ThreeSumFastWithOverFlow(nums []int) [][3]int {
	// TODO: use my sort instead <2022-05-04, genmzy> //
	sort.Ints(nums)
	res := make([][3]int, 0, 70000)
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			m := nums[i] + nums[j]
			if nums[i] < 0 && nums[j] < 0 && m > 0 {
				continue
			}
			if m > 0 {
				break
			}
			if k := binarysearch.Rank(-m, nums); k > j {
				res = append(res, [3]int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return res
}

// O(N*lgN)
func ThreeSumCntFastWithOverFlow(nums []int) int {
	// TODO: use my sort instead <2022-05-04, genmzy> //
	sort.Ints(nums)
	res := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			m := nums[i] + nums[j]
			if nums[i] < 0 && nums[j] < 0 && m > 0 {
				continue
			}
			if m > 0 {
				break
			}
			// only need to handle postive overflow
			if binarysearch.Rank(-m, nums) > j {
				res++
			}
		}
	}
	return res
}
