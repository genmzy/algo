package main

import (
	"algo/cases/ch1/1_5/dynamic_connectivity/unionfind"
	uf_extra "algo/cases/ch1/1_5/dynamic_connectivity/unionfind/extra"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// NOTE: use the b.N of benchmark for num seems wrong...
// ./bin/quickunion_rand 100000 quick_union
func main() {
	if len(os.Args) != 3 {
		log.Fatalf("usage: %s <number> <type>; type could be one of these:\n\tquick_find\n\tquick_union\n\tweight_quick_union\n\tcompress_weight_quick_union\n\theight_quick_union\n", os.Args[0])
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var uf unionfind.UnionFind
	switch os.Args[2] {
	case "quick_union":
		uf = unionfind.QuickUnionUF(num)
	case "quick_find":
		uf = unionfind.QuickFindUF(num)
	case "weight_quick_union":
		uf = unionfind.WeightQuickUnionUF(num)
	case "compress_weight_quick_union":
		uf = unionfind.CompressWeightQuickUnionUF(num)
	case "height_quick_union":
		uf = uf_extra.HeightQuickUnionUF(num)
	default:
		log.Fatalf("invalid type %s", os.Args[2])
	}
	rand.Seed(time.Now().Unix())
	var p, q int
	start := time.Now()
	for i := 0; i < num; i++ {
		uf.Union(p, q)
		p = rand.Intn(num)
		q = rand.Intn(num)
	}
	afterUnion := time.Now()
	uf.Connected(p, q)
	end := time.Now()
	log.Printf("\n\tunion time: %v;\n\tconnected time: %v\n",
		afterUnion.Sub(start), end.Sub(afterUnion))
}
