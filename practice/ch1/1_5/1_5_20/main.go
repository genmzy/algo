// for picture 1_5_8 and practice 1_5_4
package main

import (
	"algo/cases/ch1/1_5/dynamic_connectivity/unionfind"
	newUF "algo/practice/ch1/1_5/1_5_20/unionfind"
)

var data = [][2]int{
	{4, 3}, {3, 8}, {6, 5}, {9, 4}, {2, 1},
	{8, 9}, {5, 0}, {7, 2}, {6, 1}, {1, 0}, {6, 7},
}

func main() {
	dwqu := newUF.WeightQuickUnionUF()
	wqu := unionfind.WeightQuickUnionUF(10)
	// just for debug
	for _, v := range data {
		dwqu.Union(v[0], v[1])
		wqu.Union(v[0], v[1])
	}
}
