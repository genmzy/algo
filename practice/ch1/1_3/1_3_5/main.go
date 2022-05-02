package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		v, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatal(err)
		}
		s := anonymous(v)
		s.Traverse(func(i int) bool {
			fmt.Printf("%d ", i)
			return false
		})
		fmt.Println()
	}
}

func anonymous(v int) *linkstack.Stack[int] {
	s := &linkstack.Stack[int]{}
	for v > 0 {
		s.Push(v % 2)
		v = v / 2
	}
	return s
}
