package main

import (
	"algo/practice/ch1/1_3/1_3_19_to_28/linklist"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var store linklist.LinkList[byte]

func dump() {
	store.Traverse(func(c byte) bool {
		fmt.Printf("%c ", c)
		return false
	})
	fmt.Println()
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if len(s.Text()) == 0 {
			continue
		}
		units := strings.Split(s.Text(), " ")
		for _, unit := range units {
			c := unit[0]
			moveToTheFront(c)
		}
		dump()
	}
}

func moveToTheFront(c byte) {
	if store.Find(c) {
		store.Remove(c)
	}
	store.HeadAdd(c)
}
