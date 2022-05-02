package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := linkstack.Stack[string]{}
		for _, t := range strings.Split(input.Text(), " ") {
			if "-" == t {
				p := s.Pop()
				if p != nil {
					fmt.Print(*p, " ")
				}
			} else {
				s.Push(t)
			}
		}
		fmt.Println()
	}
}
