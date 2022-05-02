package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(evaluatePostfix(s.Text()))
	}
}

// 1 2 + 3 * 5 + 1 2 * +
func evaluatePostfix(express string) int {
	texts := strings.Split(express, " ")
	vals := linkstack.Stack[int]{}
	for _, t := range texts {
		switch t {
		case "+":
			fallthrough
		case "-":
			fallthrough
		case "*":
			fallthrough
		case "/":
			last := *vals.Pop()
			before := *vals.Pop()
			cal := 0
			switch t {
			case "+":
				cal = before + last
			case "-":
				cal = before - last
			case "*":
				cal = before * last
			case "/":
				cal = before / last
			}
			vals.Push(cal)
		default:
			if n, err := strconv.Atoi(t); err == nil {
				vals.Push(n)
			} else {
				log.Fatal(err)
			}
		}
	}
	return *vals.Pop()
}
