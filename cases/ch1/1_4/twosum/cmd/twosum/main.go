package main

import (
	"algo/cases/ch1/1_4/twosum"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filepath>", os.Args[0])
		os.Exit(1)
	}
	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid file: %s", err)
		os.Exit(1)
	}
	defer f.Close()
	s := make([]int, 0, 1000001)
	input := bufio.NewScanner(f)
	for input.Scan() {
		n, err := strconv.Atoi(input.Text())
		if err != nil {
			panic(err)
		}
		s = append(s, n)
	}
	start := time.Now()
	res := twosum.TwoSumBinarySearch(s)
	d := time.Now().Sub(start)
	fmt.Printf("[binary_search]total %d group numbers, cost %v\n", len(res), d)

	start = time.Now()
	res = twosum.TwoSumHash(s)
	d = time.Now().Sub(start)
	fmt.Printf("[hash_table]total %d group numbers, cost %v\n", len(res), d)
}
