package closestfarest

import (
	"math"
	"sort"
)

// practice 1.4.16
// find closest pair float values in an array
func ClosestPair(nums []float64) [2]float64 {
	res := [2]float64{}
	sort.Float64s(nums)
	l := len(nums)
	if l < 2 {
		return res
	}
	diff := math.Abs(nums[1] - nums[0])
	for i := 1; i < l; i++ {
		m := math.Abs(nums[i] - nums[i-1])
		if m < diff {
			res[0], res[1] = nums[i], nums[i-1]
		}
	}
	return res
}

// practice 1.4.17
func FarestPair(nums []float64) [2]float64 {
	sort.Float64s(nums)
	l := len(nums)
	return [...]float64{nums[l-1], nums[0]}
}

// practice 1.4.17
func FarestPairRange(nums []float64) [2]float64 {
	sort.Float64s(nums)
	max, min := nums[0], nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return [...]float64{max, min}
}
