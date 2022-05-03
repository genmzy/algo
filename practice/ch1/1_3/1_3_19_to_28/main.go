package main

import (
	"algo/practice/ch1/1_3/1_3_19_to_28/linklist"
	"fmt"
)

func dump(l *linklist.LinkList[int]) {
	l.Traverse(func(i int) bool {
		fmt.Print(i)
		fmt.Print(" ")
		return false
	})
	fmt.Println()
}

func main() {
	l := &linklist.LinkList[int]{}
	l.Add(8)
	l.Add(8)
	l.Add(8)
	for i := 0; i < 10; i++ {
		l.Add(i)
	}
	fmt.Println("max:", linklist.Max(l))

	l.InsertAfter(5, 8) // insert 8 after 8
	l.Add(8)
	fmt.Println("after add and insert:")
	dump(l)

	l.Remove(8) // remove all 8
	fmt.Println("after remove all `8':")
	dump(l)

	l.DeleteTail()
	fmt.Println("delete tail:")
	dump(l)

	l.RemoveAfter(0)
	dump(l)

	fmt.Println("now max:", linklist.MaxByRecursive(l))

	l.Reverse()
	dump(l)

	l.ReverseByRecursive()
	dump(l)
}
