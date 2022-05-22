package main

import (
	"reflect"
	"testing"
)

func TestCommonNums(t *testing.T) {
	type args struct {
		a []int
		b []int
		N int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{a: []int{1, 2, 3, 4, 5}, b: []int{4, 5, 6, 7, 8}}, want: []int{4, 5}},
		{args: args{a: []int{1, 2, 3, 4, 5}, b: []int{4, 5, 6}}, want: []int{4, 5}},
		{args: args{a: []int{1, 2, 3}, b: []int{1, 2, 3}}, want: []int{1, 2, 3}},
		{args: args{b: []int{1, 2, 3, 4, 5}, a: []int{4, 5, 6, 7, 8}}, want: []int{4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommonNums(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommonNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
