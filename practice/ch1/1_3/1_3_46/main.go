package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.OpenFile("practice/ch1/1_3/1_3_3_13_45/is_poped_from_ordered_stack.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text()[0] == '#' {
			continue
		}
		target := make([]int, 0)
		for _, v := range strings.Split(input.Text(), " ") {
			n, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			target = append(target, n)
		}
		fmt.Printf("%s is poped from order stack: %v\n", input.Text(),
			isPoppedFromOrderedStack([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, target))
	}
}

func isPoppedFromOrderedStack(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		fmt.Println("different lenth")
		os.Exit(1)
	}
	for i := 0; i < len(popped)-2; i++ {
		// big small medium
		if popped[i] > popped[i+2] && popped[i+2] > popped[i+1] {
			return false
		}
	}
	return true
}
