package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandDup(top int) int {
	store := make(map[int]bool)
	for i := 0; ; i++ {
		x := rand.Intn(top + 1)
		if store[x] {
			return i
		} else {
			store[x] = true
		}
	}
}

func CollectAll(top int) int {
	store := make(map[int]bool)
	length := 0
	for i := 0; ; i++ {
		x := rand.Intn(top + 1)
		if store[x] {
			continue
		} else {
			store[x] = true
			length++
		}
		if length == top {
			return i
		}
	}
}

func main() {
	for _, arg := range os.Args[1:] {
		i, err := strconv.Atoi(arg)
		if err != nil {
			continue
		}
		fmt.Printf("%d random numbers create duplicate in %d numbers\n", RandDup(i), i)
	}
}
