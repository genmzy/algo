package main

import (
	fibsearch "algo/practice/ch1/1_4/1_4_22"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 7, 8, 231, 3248, 32378, 39841, 328908, 4311131}
	fmt.Println(fibsearch.FibonacciRank(nums, 231))
	fmt.Println(fibsearch.FibRank(nums, 231))
}
