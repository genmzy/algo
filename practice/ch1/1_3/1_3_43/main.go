package main

import (
	"algo/cases/ch1/1_3/linkstack"
	"algo/cases/ch1/1_3/queue"
	"fmt"
	"os"
)

var filenames = &queue.Queue[string]{}
var prefix = &linkstack.Stack[string]{} // file or directory prefix

func expand(dn string) {
	entries, err := os.ReadDir(dn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read directory: %v\n", err)
		return
	}
	prefix.Push(dn)
	defer prefix.Pop()
	for _, e := range entries {
		name := *prefix.Peek() + "/" + e.Name()
		if e.IsDir() {
			expand(name)
		} else {
			filenames.Enqueue(name)
		}
	}
}

// list all files in directory recursively
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <directory>", os.Args[0])
		os.Exit(1)
	}
	dn := os.Args[1]
	expand(dn)
	filenames.Traverse(func(dn string) bool {
		fmt.Println(dn)
		return false
	})
}
