package main

import (
	"algo/cases/ch1/1_3/queue"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	q := queue.Queue[string]{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		units := strings.Split(s.Text(), " ")
		for _, unit := range units {
			q.Enqueue(unit)
		}
		buf := bytes.Buffer{}
		q.Traverse(func(s string) bool {
			buf.WriteString(s)
			buf.WriteString(" -> ")
			return false
		})
		buf.WriteString("nil")
		fmt.Println("queue dump:", buf.String())
		fmt.Println("size of queue:", q.Size())
	}

	for q.Size() != 0 {
		s := *q.Dequeue()
		fmt.Println("dequeue ", s)
	}
}
