package main

import (
	"algo/practice/ch1/1_4/1_4_18_19/localminimum"
	"fmt"
)

func main() {
	matrix()
}

func matrix() {
	matrix := [][]int{
		{9, 3, 5, 2, 4, 8, 9},
		{7, 2, 5, 1, 4, 0, 3},
		{9, 8, 9, 3, 2, 4, 8},
		{7, 6, 3, 1, 3, 2, 3},
		{9, 0, 6, 0, 4, 6, 4},
		{8, 9, 8, 0, 5, 3, 0},
		{2, 1, 2, 1, 1, 1, 1},
	}
	i, j := localminimum.MatrixLocalMimium(matrix)
	fmt.Println(i, j)

	i, j = localminimum.MatrixLocalMinimumFaster(matrix)
	fmt.Println(i, j)
}
