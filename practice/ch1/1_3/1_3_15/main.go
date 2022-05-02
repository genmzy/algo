package main

import (
	"algo/cases/ch1/1_3/queue"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <k>", os.Args[0])
	}
	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("k must be a number")
	}
	q := &queue.Queue[string]{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		units := strings.Split(s.Text(), " ")
		if len(units) < k {
			fmt.Fprintf(os.Stderr, "length of all strings must be greater than %d\n", k)
			continue
		}
		for _, t := range units {
			q.Enqueue(t)
		}
		log.Printf("the last %d th is: %s\n", k, *getFromDescKth(q, k))
	}
}

func getFromDescKth[T any](q *queue.Queue[T], k int) *T {
	i := 0
	var v *T
	q.Traverse(func(t T) bool {
		i++
		if i == q.Size()-k+1 {
			v = &t
			return true
		}
		return false
	})
	return v
}
