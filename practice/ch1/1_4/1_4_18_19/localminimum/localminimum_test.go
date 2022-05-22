package localminimum

import (
	"testing"
)

func TestArrayLocalMinimum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{1, 0, 3, 5, 4}}, want: 1},
		{args: args{nums: []int{2, 1, 3, 5, 4}}, want: 1},
		{args: args{nums: []int{3, 1, 4, 5, 2, 6}}, want: 1},
		{args: args{nums: []int{2, 1, 2, 1, 1, 1, 1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayLocalMinimum(tt.args.nums); got != tt.want {
				t.Errorf("LocalMinimum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrixLocalMinimumFaster(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MatrixLocalMinimumFaster(tt.args.matrix)
			if got != tt.want {
				t.Errorf("MatrixLocalMinimumFaster() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MatrixLocalMinimumFaster() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
