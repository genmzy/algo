package main

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkDoubleAdjustFind(b *testing.B) {
	rand.Seed(time.Now().Unix())
	step := rand.Intn(50)
	total := 100000
	for i := 0; i < b.N; i++ {
		num := rand.Intn(step)
		nums := make([]int, 0, total+1)
		n := rand.Intn(total * step)
		for i := 0; i < total; i++ {
			nums = append(nums, i*step+num)
		}
		//b.Run(DoubleAdjustableFind(nums, n))
		b.Run("", func(b *testing.B) {
			DoubleAdjustableFind(nums, n)
		})
	}
}
