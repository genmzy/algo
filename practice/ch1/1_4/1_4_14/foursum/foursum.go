package foursum

import (
	"algo/cases/ch1/1_2/intset"
	"sort"
)

func FourSum(nums []int) [][4]int {
	n := len(nums)
	res := make([][4]int, 0, n)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if i != 0 && nums[i] != nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < n; k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				sum := nums[i] + nums[j] + nums[k]
				if sum > 0 {
					break
				}
				if m := intset.IntSet(nums).Rank(-sum); m > k {
					res = append(res, [...]int{nums[i], nums[j], nums[k], nums[m]})
				}
			}
		}
	}
	return res
}
