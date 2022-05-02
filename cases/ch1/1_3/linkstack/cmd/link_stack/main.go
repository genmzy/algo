package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	s := linkstack.Stack[int]{}
	buf := bytes.Buffer{}
	for i := 0; i <= 10; i++ {
		s.Push(i)
	}
	s.Traverse(func(v int) bool {
		buf.WriteString(strconv.Itoa(v))
		buf.WriteString(" -> ")
		return false
	})
	buf.WriteString("nil")
	fmt.Println(buf.String())
}
