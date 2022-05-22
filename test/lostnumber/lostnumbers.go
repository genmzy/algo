package main

func LostNumbers(nums []int, max int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	res := make([]int, 0, nums[n-1])
	k := 0 // index
	for i := 1; i <= max; i++ {
		res = append(res, i)
		if k == n {
			continue
		} else {
			if i == nums[k] {
				res = res[:len(res)-1]
				k++
			}
		}
	}
	return res
}
