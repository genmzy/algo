package main

import (
	"algo/cases/ch1/1_4/stopwatch"
	"algo/cases/ch1/1_4/threesum"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <number>\n", os.Args[0])
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid number %s\n", os.Args[1])
		os.Exit(1)
	}
	rand.Seed(time.Now().Unix())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		k := rand.Intn(1000000*2) - 1000000
		nums = append(nums, k)
	}
	timer := stopwatch.NewStopWatch()
	cnt := threesum.ThreeSumCntFast(nums)
	fmt.Println(cnt, "triples", timer.ElapsedTime())
}
