// ErdosRenyi case
package main

import (
	"algo/practice/ch1/1_5/1_5_17/randconnect"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s <num>", os.Args[0])
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	log.Println("count size:", randconnect.Count(n))
	return
}
