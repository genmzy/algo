package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(infixToPostfix(s.Text()))
	}
}

func infixToPostfix(express string) string {
	texts := strings.Split(express, " ")
	ops := linkstack.Stack[string]{}
	vals := linkstack.Stack[string]{}

	for _, t := range texts {
		switch t {
		case "(":
		case "+":
			fallthrough
		case "-":
			fallthrough
		case "*":
			fallthrough
		case "/":
			ops.Push(t)
		case ")":
			v := *vals.Pop()
			buf := bytes.Buffer{}
			buf.WriteString(*vals.Pop())
			buf.WriteByte(' ')
			buf.WriteString(v)
			buf.WriteByte(' ')
			buf.WriteString(*ops.Pop())
			vals.Push(buf.String())
		default:
			vals.Push(t)
		}
	}
	return *vals.Pop()
}
