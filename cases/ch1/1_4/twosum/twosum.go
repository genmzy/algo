package twosum

import (
	"algo/cases/ch1/1_1/binarysearch"
	"sort"
)

// O(n*n)
func TwoSum(nums []int) [][2]int {
	res := make([][2]int, 0, 300000)
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] > nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j] == 0 {
				res = append(res, [...]int{nums[i], nums[j]})
			}
		}
	}
	return res
}

// O(n*n)
func TwoSumCnt(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] > nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j] == 0 {
				res++
			}
		}
	}
	return res
}

// O(N*lgN)
func TwoSumBinarySearch(nums []int) [][2]int {
	// TODO: use my sort instead <2022-05-04, genmzy> //
	sort.Ints(nums)
	res := make([][2]int, 0, 300000) // 240k <= result(already knows) <= 260k
	n := len(nums)
	for i := 0; i < n; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		// TODO: figure out why this is wrong <2022-05-04, genmzy> //
		// if j := binarysearch.Rank(-nums[i], nums[i:]); j > i {
		if j := binarysearch.Rank(-nums[i], nums); j > i {
			res = append(res, [2]int{nums[i], nums[j]})
		}
	}
	return res
}

// O(N*lgN)
func TwoSumCntFast(nums []int) int {
	// TODO: use my sort instead <2022-05-04, genmzy> //
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i+1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		if binarysearch.Rank(-nums[i], nums) > i {
			res++
		}
	}
	return res
}

func TwoSumHash(nums []int) [][2]int {
	hash := make(map[int]int) // num -> idx
	n := len(nums)
	res := make([][2]int, 0, n)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		k := nums[i]
		if k < 0 {
			hash[k] = i
		} else {
			if _, ok := hash[-k]; ok {
				res = append(res, [...]int{k, -k})
			}
		}
	}
	return res
}
