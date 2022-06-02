package main

import "algo/cases/ch1/1_5/dynamic_connectivity/unionfind"

var data = [][2]int{
	{8, 4}, {4, 5}, {9, 5}, {5, 6}, {2, 3},
	{7, 3}, {3, 1}, {6, 1}, {0, 1},
}

func main() {
	uf := unionfind.QuickUnionUF(10)
	for _, v := range data {
		uf.Union(v[0], v[1])
	}
	return
}
