package fractionfind

import (
	"math"
)

/**
 * 1.4.23 Binary search for a fraction.
 * Devise a method that uses a logarithmic number of queries of the form:
 * Is the number less than x?
 *
 * to find a rational number p/q such that 0 < p < q < N.
 *
 * Hint: Two fractions with denominators less than N cannot differ by more than 1/(N^2).
 *
 * Note: Hint is incorrect. Opposite is true.
 *  1/p - 1/q >= 1/n^2 ( n > p && n > q)
 *  (q-p)/pq >= 1/n^2
 *
 * Note2: Question is unclear... p/q is always a fraction number. What is it asking us to find?
 *
 */
func FractionBinFind(nums []float64, k float64) int {
	n := len(nums)
	l, r := 0, n-1
	var threshold float64 = 1.0 / (float64(n) * float64(n))
	for l <= r {
		mid := l + (r-l)>>1
		if math.Abs(nums[mid]-k) <= threshold {
			return mid
		}
		if nums[mid] < k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
