package main

import (
	"algo/cases/ch1/1_5/dynamic_connectivity/unionfind"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s <num>\n", os.Args[0])
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	qu := unionfind.QuickUnionUF(n)
	start := time.Now()
	for i := 0; i < n-1; i++ {
		qu.Union(0, i+1) // NOTE: union 0 and n makes always bad
	}
	afterUnion := time.Now()
	qu.Connected(0, n)
	end := time.Now()
	log.Printf("quick-union UF:\n\tunion time: %v;\n\tconnection time: %v\n",
		afterUnion.Sub(start), end.Sub(afterUnion))

	qf := unionfind.QuickFindUF(n)
	start = time.Now()
	for i := 0; i < n-1; i++ {
		qf.Union(i, i+1)
	}
	afterUnion = time.Now()
	qf.Connected(0, n)
	end = time.Now()
	log.Printf("quick-find UF:\n\tunion time: %v;\n\tconnection time: %v\n",
		afterUnion.Sub(start), end.Sub(afterUnion))
}
