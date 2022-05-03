package main

import (
	"algo/cases/ch1/1_3/queue"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dump(q *queue.Queue[string]) {
	fmt.Print("dumping ")
	q.Traverse(func(s string) bool {
		fmt.Printf("%s ", s)
		return false
	})
	fmt.Println("\nlenth:", q.Size())
}

func main() {
	f, err := os.OpenFile("cases/ch1/1_3/queue/tobe.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	q := &queue.Queue[string]{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		units := strings.Split(s.Text(), " ")
		for _, unit := range units {
			q.Enqueue(unit)
		}
		dump(q)
	}

	with := &queue.Queue[string]{}
	for i := 100; i > 90; i-- {
		with.Enqueue(strconv.Itoa(i))
	}
	dump(with)

	q.Catenation(with)
	dump(q)

	for q.Size() != 0 {
		s := *q.Dequeue()
		fmt.Println("dequeue ", s)
	}
}
