package unionfind

import "algo/cases/ch1/1_5/dynamic_connectivity/unionfind"

func weightQuickUnion() *weightQuickUnionUF {
	qu := &weightQuickUnionUF{
		roots:     make([]int, 0),
		treeSizes: make([]int, 0),
		count:     0,
	}
	for i := 0; i < qu.count; i++ {
		qu.roots = append(qu.roots, i)
	}
	return qu
}

// improved union find algorithm: make the
func WeightQuickUnionUF() unionfind.UnionFind {
	return weightQuickUnion()
}

func WeightQuickUnionUFWithVisit() unionfind.UnionFindWithVisit {
	return weightQuickUnion()
}

type weightQuickUnionUF struct {
	roots     []int // contact id => root contact id
	treeSizes []int // root contact id => size of tree of root contact id
	count     int   // component count
}

func max(p, q int) int {
	if p > q {
		return p
	}
	return q
}

func (qu *weightQuickUnionUF) Union(p, q int) {
	if l := len(qu.roots); p >= l || q >= l {
		old := len(qu.roots)
		up := max(p, q)
		for i := old; i <= up; i++ {
			qu.roots = append(qu.roots, i)
		}
		for i := old; i <= up; i++ {
			qu.treeSizes = append(qu.treeSizes, 1)
		}
		qu.count += (up - old)
	}
	pRoot := qu.Find(p)
	qRoot := qu.Find(q)
	if pRoot == qRoot {
		return
	}
	if qu.treeSizes[pRoot] < qu.treeSizes[qRoot] {
		qu.roots[pRoot] = qRoot
		qu.treeSizes[qRoot] += qu.treeSizes[pRoot]
	} else {
		qu.roots[qRoot] = pRoot
		qu.treeSizes[pRoot] += qu.treeSizes[qRoot]
	}
	qu.count--
}

func (qu *weightQuickUnionUF) UnionWithVisit(p, q int) int {
	if l := len(qu.roots); p >= l || q >= l {
		return 0
	}
	v := 0
	pRoot, k := qu.FindWithVisit(p)
	v += k
	qRoot, k := qu.FindWithVisit(q)
	v += k
	if pRoot == qRoot {
		return v
	}
	// NOTE: always merge the smaller tree to the bigger tree
	// makes the depth of the tree always smaller than lgN
	if qu.treeSizes[pRoot] < qu.treeSizes[qRoot] {
		qu.roots[pRoot] = qRoot
		qu.treeSizes[qRoot] += qu.treeSizes[pRoot]
	} else {
		qu.roots[qRoot] = pRoot
		qu.treeSizes[pRoot] += qu.treeSizes[qRoot]
	}
	v += 3
	qu.count--
	return v
}

// find the root of p
func (qu *weightQuickUnionUF) Find(p int) int {
	if l := len(qu.roots); p >= l {
		return -1
	}
	for p != qu.roots[p] {
		p = qu.roots[p]
	}
	return p
}

func (qu *weightQuickUnionUF) FindWithVisit(p int) (int, int) {
	if l := len(qu.roots); p >= l {
		return -1, 0
	}
	v := 0
	for p != qu.roots[p] {
		p = qu.roots[p]
		v += 2
	}
	v++
	return p, v
}

// if contact p and contact q in the same weight, return true
func (qu *weightQuickUnionUF) Connected(p, q int) bool {
	if l := len(qu.roots); p >= l || q >= l {
		return false
	}
	return qu.Find(p) == qu.Find(q)
}

func (qu *weightQuickUnionUF) Count() int {
	return qu.count
}
