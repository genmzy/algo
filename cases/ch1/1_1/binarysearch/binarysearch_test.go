package binarysearch

import "testing"

func TestRank(t *testing.T) {
	type args struct {
		n    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{n: 5, nums: []int{2, 4, 5, 7, 10, 15}}, want: 2},
		{args: args{n: 7, nums: []int{2, 4, 5, 7, 10, 15}}, want: 3},
		{args: args{n: 3, nums: []int{2, 4, 5, 7, 10, 15}}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rank(tt.args.n, tt.args.nums); got != tt.want {
				t.Errorf("Rank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMinIdx(t *testing.T) {
	type args struct {
		n    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// 0 + (4-0)/2 = 2
		{args: args{n: 5, nums: []int{2, 4, 5, 5, 5, 5, 7, 10, 15}}, want: 2},
		{args: args{n: 3, nums: []int{2, 4, 5, 5, 5, 5, 7, 10, 15}}, want: -1},
		{args: args{n: 7, nums: []int{2, 4, 5, 5, 5, 5, 7, 10, 15}}, want: 6},
		{args: args{n: 99, nums: []int{7, 19, 28, 76, 76, 99, 99, 99, 99}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMinIdx(tt.args.n, tt.args.nums); got != tt.want {
				t.Errorf("RankMini() = %v, want %v", got, tt.want)
			}
		})
	}
}
