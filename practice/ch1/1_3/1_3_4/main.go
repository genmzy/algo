package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args[1:] {
		fmt.Printf("argv[%d]: %v\n", i, parentheses([]byte(v)))
	}
}

// m is left, n is right
func isMatch(m byte, n byte) bool {
	switch n {
	case '[':
		return m == ']'
	case '{':
		return m == '}'
	case '<':
		return m == '>'
	case '(':
		return m == ')'
	case '\'':
		fallthrough
	case '"':
		return m == n
	default:
		return false
	}
}

func parentheses(v []byte) bool {
	s := linkstack.Stack[byte]{}
	for _, c := range v {
		if s.Empty() {
			s.Push(c)
			continue
		}
		if !isMatch(c, *s.Top()) {
			s.Push(c)
		} else {
			s.Pop()
		}
	}
	return s.Empty()
}
