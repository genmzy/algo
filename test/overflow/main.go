package main

import "fmt"

func main() {
	var dummy int64 = ^1
	a := int(uint64(dummy) >> 1)
	fmt.Printf("a: %d\n", a)
}
