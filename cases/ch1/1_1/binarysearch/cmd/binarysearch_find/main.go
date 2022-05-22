package main

import (
	"algo/cases/ch1/1_1/binarysearch"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// read from database
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <filepath>\n", os.Args[0])
		os.Exit(1)
	}
	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid file: %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	defer f.Close()
	// load databases data to memory
	list := make([]int, 0)
	r := bufio.NewScanner(f)
	for r.Scan() {
		n, err := strconv.Atoi(r.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid number: %s\n", r.Text())
			os.Exit(1)
		}
		list = append(list, n)
	}
	// interactive with scanner
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s is not a number\n", s.Text())
			continue
		}
		binarysearch.Rank(i, list)
	}
}
