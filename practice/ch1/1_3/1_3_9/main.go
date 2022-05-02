package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		t := input.Text()
		s := addLeftBrackets(t)
		fmt.Println("result is:", s)
		//s = addLeftBracketsErr(input.Text())
		//fmt.Println("old result is:", s)
	}
}

// 1 + 2 * 3 - 4 ) * 5 - 6 ) ) )
func addLeftBrackets(express string) string {
	texts := strings.Split(express, " ")
	vals := &linkstack.Stack[string]{}
	ops := &linkstack.Stack[string]{}
	for _, t := range texts {
		switch t {
		case "+":
			fallthrough
		case "-":
			fallthrough
		case "*":
			fallthrough
		case "/":
			ops.Push(t)
		case ")":
			op := *ops.Pop()
			v := *vals.Pop()
			buf := bytes.Buffer{}
			buf.WriteString("( ")
			buf.WriteString(*vals.Pop())
			buf.WriteByte(' ')
			buf.WriteString(op)
			buf.WriteByte(' ')
			buf.WriteString(v)
			buf.WriteString(" )")
			// regard original express as a val
			vals.Push(buf.String())
		default:
			vals.Push(t)
		}
	}
	return *vals.Pop()
}

func addLeftBracketsErr(express string) string {
	preNum := false     // previous number already exist
	signUndisposed := 0 // add '(' regards as disposed, eg: ` 3 - 4 ) * ( 5 - 6 )` signUndisposed is 2(`*` and '-')
	s := &linkstack.Stack[string]{}
	txts := strings.Split(express, " ")
	for i := len(txts) - 1; i >= 0; i-- {
		txt := txts[i]
		s.Push(txt)
		switch txt {
		case "+":
			fallthrough
		case "-":
			fallthrough
		case "*":
			fallthrough
		case "/":
			signUndisposed++
		default:
			if _, err := strconv.ParseFloat(txt, 64); err == nil {
				if preNum {
					for ; signUndisposed >= 0; signUndisposed-- {
						s.Push("(")
					}
					preNum = false
				} else {
					preNum = true
				}
			}

		}
	}
	buf := bytes.Buffer{}
	s.Traverse(func(s string) bool {
		buf.WriteString(s)
		buf.WriteByte(' ')
		return false
	})
	return buf.String()
}
