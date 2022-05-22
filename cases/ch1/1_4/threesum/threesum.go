package threesum

import (
	"algo/cases/ch1/1_1/binarysearch"
	"sort"
)

// give a sum
func ThreeSum(nums []int) [][3]int {
	res := make([][3]int, 0)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] != nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[k] == nums[k-1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, [...]int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return res
}

func ThreeSumCnt(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[k] == nums[k-1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					res++
				}
			}
		}
	}
	return res
}

func ThreeSumBinarySearch(nums []int) [][3]int {
	// TODO: use my sort instead <2022-05-04, genmzy> //
	sort.Ints(nums)
	res := make([][3]int, 0, 70000)
	n := len(nums)
	for i := 0; i < n; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if k := binarysearch.Rank(-nums[i]-nums[j], nums); k > j {
				res = append(res, [3]int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return res
}

// O(N*lgN)
func ThreeSumCntFast(nums []int) int {
	// TODO: use my sort instead <2022-05-04, genmzy> //
	sort.Ints(nums)
	res := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			m := nums[i] + nums[j]
			if m > 0 {
				break
			}
			if binarysearch.Rank(-m, nums) > j {
				res++
			}
		}
	}
	return res
}

func ThreeSumDoublePointer(nums []int) [][3]int {
	n := len(nums)
	sort.Ints(nums)
	if nums[0] > 0 || nums[n-1] < 0 || n < 3 {
		return [][3]int{}
	}
	if nums[0] == 0 && nums[n-1] == 0 {
		return [][3]int{[...]int{0, 0, 0}}
	}
	res := make([][3]int, 0, n)
	i, j, k := 0, 0, 0
	for i = 0; i < n-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		j = i + 1
		k = n - 1
		for j < k {
			m := nums[i] + nums[j] + nums[k]
			if m > 0 {
				k--
			} else if m < 0 {
				j++
			} else {
				res = append(res, [...]int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return res
}
