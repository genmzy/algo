package intset

import "testing"

var set = []int{1, 3, 3, 3, 3, 4, 5, 5, 6, 7, 7, 8, 8, 8, 10, 11, 12, 13, 14, 15}

func TestIntSet_Contains(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		s    IntSet
		args args
		want bool
	}{
		{s: set, args: args{3}, want: true},
		{s: set, args: args{2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.n); got != tt.want {
				t.Errorf("IntSet.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_HowMany(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		s    IntSet
		args args
		want int
	}{
		{s: set, args: args{2}, want: 0},
		{s: set, args: args{3}, want: 4},
		{s: set, args: args{5}, want: 2},
		{s: set, args: args{8}, want: 3},
		{s: set, args: args{6}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.HowMany(tt.args.n); got != tt.want {
				t.Errorf("IntSet.HowMany() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_Rank(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		s    IntSet
		args args
		want int
	}{
		{s: set, args: args{6}, want: 8},
		{s: set, args: args{4}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Rank(tt.args.n); got != tt.want {
				t.Errorf("IntSet.Rank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSet_MiniIdx(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		s    IntSet
		args args
		want int
	}{
		{s: set, args: args{6}, want: 8},
		{s: set, args: args{1}, want: 0},
		{s: set, args: args{3}, want: 1},
		{s: set, args: args{5}, want: 6},
		{s: set, args: args{7}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.MiniIdx(tt.args.n); got != tt.want {
				t.Errorf("IntSet.RankMini() = %v, want %v", got, tt.want)
			}
		})
	}
}
