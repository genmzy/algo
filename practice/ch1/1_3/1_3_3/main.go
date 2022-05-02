package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		target := make([]int, 0)
		for _, v := range strings.Split(input.Text(), " ") {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			target = append(target, n)
		}
		fmt.Printf("target array is poped from order stack: %v\n",
			isPoppedFromOrderedStack(target))
	}
}

// TODO: finsh this later <25-04-22, genmzy> //
func isPoppedFromOrderedStack(target []int) bool {
	pre := -1
	prepre := -1
	for i, v := range target {
		if i < 2 {
			continue
		}
		if i == 2 {
			prepre = target[0]
			pre = target[1]
		}
		if prepre >= v && v >= pre {
			return false
		}
		prepre = pre
		pre = v
	}
	return true
}
