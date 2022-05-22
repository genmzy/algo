package main

import (
	"algo/practice/ch1/1_4/1_4_2/threesum"
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
	n := threesum.ThreeSumCntFastWithOverFlow(s)
	d := time.Now().Sub(start)
	fmt.Printf("total %d types, cost %v\n", n, d)
}
