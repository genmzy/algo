package threesum

import "testing"

func BenchmarkThreeSumCntWithOverFlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThreeSumFastWithOverFlow([]int{-1, -2 - 3 - 4, 4, 3, 2, 1, 0})
	}
}
