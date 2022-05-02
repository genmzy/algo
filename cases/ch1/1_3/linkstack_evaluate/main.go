package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		units := strings.Split(s.Text(), " ")
		ops := linkstack.Stack[string]{}
		vals := &linkstack.Stack[float64]{}
		for _, unit := range units {
			switch unit {
			case "(":
			case "+":
				fallthrough
			case "-":
				fallthrough
			case "*":
				fallthrough
			case "/":
				fallthrough
			case "sqrt":
				ops.Push(unit)
			case ")":
				op := ops.Pop()
				v := *vals.Pop()
				switch *op {
				case "+":
					v = *vals.Pop() + v
				case "-":
					v = *vals.Pop() - v
				case "*":
					v = *vals.Pop() * v
				case "/":
					v = *vals.Pop() / v
				case "sqrt":
					v = math.Sqrt(v)
				}
				vals.Push(v)
			default:
				num, err := strconv.ParseFloat(unit, 64)
				if err != nil {
					log.Fatal(err)
				}
				vals.Push(num)
			}
		}
		log.Println("result:", *vals.Pop())
	}
}
