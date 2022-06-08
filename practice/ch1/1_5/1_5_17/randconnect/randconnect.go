package randconnect

import (
	"algo/cases/ch1/1_5/dynamic_connectivity/unionfind"
	"math/rand"
	"time"
)

func Count(n int) int {
	rand.Seed(time.Now().Unix())
	uf := unionfind.WeightQuickUnionUF(n)
	c := 0
	for {
		if uf.Count() <= 1 && c > 1 {
			break
		}
		p := rand.Intn(n)
		q := rand.Intn(n)
		uf.Union(p, q)
		c++
	}
	return c
}
