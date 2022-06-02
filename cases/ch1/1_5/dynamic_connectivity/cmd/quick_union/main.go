package main

import (
	"algo/cases/ch1/1_5/dynamic_connectivity/unionfind"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatalf("usage: %s <num> <want_1> <want_2> <file>, got %d args\n", os.Args[0], len(os.Args))
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	qu := unionfind.QuickUnionUF(n)
	f, err := os.OpenFile(os.Args[4], os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		sl := strings.Split(s.Text(), " ")
		if len(sl) != 2 {
			log.Fatal("invalid lenth of array")
		}
		p, err := strconv.Atoi(sl[0])
		q, err := strconv.Atoi(sl[1])
		if err != nil {
			log.Fatal(err)
		}
		qu.Union(p, q)
	}

	p, err := strconv.Atoi(os.Args[2])
	q, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d and %d: %v\n", p, q, qu.Connected(p, q))
}
