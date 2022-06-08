// RandomGrid case
package main

import (
	randomgrid "algo/practice/ch1/1_5/1_5_18/random_grid"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s <number>", os.Args[0])
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	v := randomgrid.Generate(n)
	log.Println("lenth:", len(v), "\n", v)
}
