package main

import (
	"algo/cases/ch1/1_5/dynamic_connectivity/unionfind"
	"fmt"
	"log"
)

var input = [][2]int{
	{9, 0}, {3, 4}, {5, 8}, {7, 2},
	{2, 1}, {5, 7}, {0, 3}, {4, 2},
}

func main() {
	exec1_5_1()
	fmt.Println("===================================================")
	exec1_5_2()
	fmt.Println("===================================================")
	exec1_5_3()
	fmt.Println("===================================================")
	exec1_5_4()
}

func exec1_5_1() {
	uf := unionfind.QuickFindUFWithVisit(8)
	log.Println("QuickFind:")
	for _, v := range input {
		k := uf.UnionWithVisit(v[0], v[1])
		fmt.Printf("\tunion %d and %d visit data array %d times\n",
			v[0], v[1], k)
	}
}

func exec1_5_2() {
	uf := unionfind.QuickUnionUFWithVisit(8)
	log.Println("QuickUnion:")
	for _, v := range input {
		k := uf.UnionWithVisit(v[0], v[1])
		fmt.Printf("\tunion %d and %d visit data array %d times\n",
			v[0], v[1], k)
	}
}

func exec1_5_3() {
	uf := unionfind.WeightQuickUnionUFWithVisit(8)
	log.Println("WeightQuickUnion:")
	for _, v := range input {
		k := uf.UnionWithVisit(v[0], v[1])
		fmt.Printf("\tunion %d and %d visit data array %d times\n",
			v[0], v[1], k)
	}
}

func exec1_5_4() {
	uf := unionfind.CompressWeightQuickUnionUFWithVisit(8)
	log.Println("CompressWeightQuickUnion:")
	for _, v := range input {
		k := uf.UnionWithVisit(v[0], v[1])
		fmt.Printf("\tunion %d and %d visit data array %d times\n",
			v[0], v[1], k)
	}
}
