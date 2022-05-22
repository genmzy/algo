package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <N>\n", os.Args[0])
		os.Exit(1)
	}
	cnt, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("invalid number", os.Args[1])
		os.Exit(1)
	}
	k := 0
	for n := cnt; n > 0; n /= 2 {
		for i := 0; i < n; i++ {
			k++
		}
	}
	fmt.Println("k is", k)
}
