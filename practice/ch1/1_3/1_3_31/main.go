package main

import (
	"algo/practice/ch1/1_3/1_3_31/bilinklist"
	"fmt"
)

func dump(l bilinklist.BiLinkList[int]) {
	fmt.Print("dumping: ")
	l.Traverse(func(v int) bool {
		fmt.Print(v, " ")
		return false
	}, false)
	fmt.Println()
}

func main() {
	l := bilinklist.BiLinkList[int]{}
	for i := 0; i < 10; i++ {
		l.HeadAdd(i)
	}
	dump(l)

	fmt.Print("deleting: ")
	for i := 0; i < 10; i++ {
		i := *l.TailDelete()
		fmt.Print(i, " ")
	}
	fmt.Println()
	dump(l)

	l.HeadAdd(1)
	l.AddAfter(1, 2)
	dump(l)

	l.AddAfter(1, 1)
	dump(l)
}
