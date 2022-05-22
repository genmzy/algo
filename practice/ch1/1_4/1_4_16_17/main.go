package main

import (
	closestfarest "algo/practice/ch1/1_4/1_4_16_17/closest_farest"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func randFloat64(n int) []float64 {
	ret := make([]float64, 0, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		x := rand.Float64()
		ret = append(ret, x)
	}
	return ret
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <n>", os.Args[0])
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid number: %s", os.Args[1])
		os.Exit(1)
	}
	s := randFloat64(n)

	start := time.Now()
	res := closestfarest.ClosestPair(s)
	d := time.Now().Sub(start)
	fmt.Printf("closest group: %v, cost %v\n", res, d)

	start = time.Now()
	res = closestfarest.FarestPair(s)
	d = time.Now().Sub(start)
	fmt.Printf("farest group: %v, cost %v\n", res, d)

	start = time.Now()
	res = closestfarest.FarestPairRange(s)
	d = time.Now().Sub(start)
	fmt.Printf("farest group: %v, cost %v\n", res, d)
}
