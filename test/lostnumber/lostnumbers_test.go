package main

import (
	"reflect"
	"testing"
)

func TestLostNumbers(t *testing.T) {
	type args struct {
		nums []int
		max  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums: []int{1, 3, 7}, max: 8}, want: []int{2, 4, 5, 6, 8}},
		{args: args{nums: []int{2, 4}, max: 8}, want: []int{1, 3, 5, 6, 7, 8}},
		{args: args{nums: []int{3, 8, 9}, max: 9}, want: []int{1, 2, 4, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LostNumbers(tt.args.nums, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LostNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
