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

func TimeTrail(n int) (int, time.Duration) {
	max := 1000000
	nums := make([]int, 0, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		k := rand.Intn(2*max) - max
		nums = append(nums, k)
	}
	sw := stopwatch.NewStopWatch()
	res := threesum.ThreeSumCntFast(nums)
	return res, sw.ElapsedTime()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <times>\n", os.Args[0])
		os.Exit(1)
	}
	times, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid number %s\n", os.Args[1])
		os.Exit(1)
	}
	for n := 250; ; n += n {
		if times <= 0 {
			break
		}
		cnt, d := TimeTrail(n)
		fmt.Println(cnt, "triples", d)
		times--
	}
}
