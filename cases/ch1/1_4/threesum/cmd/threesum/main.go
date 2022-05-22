package main

import (
	"algo/cases/ch1/1_4/threesum"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
	}
	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid file: %s", err)
		os.Exit(1)
	}
	defer f.Close()
	s := make([]int, 0, 10001)
	input := bufio.NewScanner(f)
	for input.Scan() {
		n, err := strconv.Atoi(input.Text())
		if err != nil {
			panic(err)
		}
		s = append(s, n)
	}
	start := time.Now()
	res := threesum.ThreeSumBinarySearch(s)
	d := time.Now().Sub(start)
	fmt.Printf("[binary_search]total %d group numbers, cost %v\n", len(res), d)

	start = time.Now()
	res = threesum.ThreeSumDoublePointer(s)
	d = time.Now().Sub(start)
	fmt.Printf("[double_pointer]total %d group numbers, cost %v\n", len(res), d)
}
