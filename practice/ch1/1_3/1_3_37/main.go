package main

import (
	"algo/cases/ch1/1_3/queue"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

/**
 * 1.3.37 Josephus problem.
 * In the Josephus problem from antiquity, N people are in dire straits and agree to the following
 * strategy to reduce the population. They arrange themselves in a circle (at positions numbered
 * from 0 to N–1) and proceed around the circle, eliminating every Mth person until only one person
 * is left. Legend has it that Josephus figured out where to sit to avoid being eliminated. Write a
 * Queue client Josephus that takes N and M from the command line and prints out the order in which
 * people are eliminated (and thus would show Josephus where to sit in the circle).
 *
 *     % java Josephus 7 2
 *     1 3 5 0 4 2 6
 */
func main() {
	q := &queue.Queue[int]{}
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <N> <M>(0<M<=N)", os.Args[0])
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid N: %s", os.Args[1])
		os.Exit(1)
	}
	m, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid M: %s\n", os.Args[2])
		os.Exit(1)
	}
	if m < 0 || n < 0 || n < m {
		fmt.Fprintf(os.Stderr, "0<M<=N\n")
		os.Exit(1)
	}
	for i := 0; i < n; i++ {
		q.Enqueue(i)
	}
	buf := bytes.Buffer{}
	for i := 1; ; i++ {
		v := q.Dequeue()
		if v == nil {
			break
		}
		if i%m > 0 {
			q.Enqueue(*v)
		} else {
			buf.WriteString(strconv.Itoa(*v))
			buf.WriteByte(' ')
		}
	}
	fmt.Println(buf.String())
}
