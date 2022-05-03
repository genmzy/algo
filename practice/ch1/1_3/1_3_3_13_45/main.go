package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("1_3_3_13:")
	q1_3_3_13()
	fmt.Println("1_3_3_45:")
	q1_3_3_45()
}

func q1_3_3_45() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text()[0] == '#' {
			return
		}
		target := make([]int, 0)
		max := -1
		for _, v := range strings.Split(input.Text(), " ") {
			n, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if n < 0 {
				fmt.Printf("invalid number: %d\n", n)
				os.Exit(1)
			}
			if n > max {
				max = n
			}
			target = append(target, n)
		}
		pushed := make([]int, 0)
		for i := 0; i <= max; i++ {
			pushed = append(pushed, i)
		}
		if len(pushed) != len(target) {
			fmt.Println(pushed, target)
			fmt.Println("invalid popped stack")
			os.Exit(1)
		}
		fmt.Printf("%s is popped from order stack: %v\n", input.Text(),
			isPoppedFromOrderedStack(pushed, target))
	}
}

func q1_3_3_13() {
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
		return false
	}
	s := linkstack.Stack[int]{}
	j := 0
	for i := 0; i < len(pushed); i++ {
		s.Push(pushed[i])
		for !s.Empty() && *s.Peek() == popped[j] {
			s.Pop()
			j++
		}
	}
	if s.Empty() {
		return true
	}
	return false
}
