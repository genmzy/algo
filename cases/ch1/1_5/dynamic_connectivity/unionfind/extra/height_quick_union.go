package uf_extra

import "algo/cases/ch1/1_5/dynamic_connectivity/unionfind"

func HeightQuickUnionUF(n int) unionfind.UnionFind {
	qu := &heightQuickUnionUF{
		roots: make([]int, 0, n),
		count: n,
	}
	// why initialize components like this?
	//    to store contacts components as all contacts are not
	//    in one component
	for i := 0; i < qu.count; i++ {
		qu.roots = append(qu.roots, i)
	}
	qu.treeDepths = make([]int, n, n) // already set to 0
	return qu
}

type heightQuickUnionUF struct {
	roots      []int // contact id => root contact id
	treeDepths []int // root contact id => size of tree of root contact id
	count      int   // component count
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func (qu *heightQuickUnionUF) Union(p, q int) {
	if l := len(qu.roots); p >= l || q >= l {
		return
	}
	pRoot := qu.Find(p)
	qRoot := qu.Find(q)
	if pRoot == qRoot {
		return
	}
	// NOTE: always merge the smaller tree to the bigger tree
	// makes the depth of the tree always smaller than lgN
	if qu.treeDepths[pRoot] < qu.treeDepths[qRoot] {
		qu.roots[pRoot] = qRoot
	} else if qu.treeDepths[pRoot] > qu.treeDepths[qRoot] {
		qu.roots[qRoot] = pRoot
	} else {
		qu.roots[qRoot] = pRoot
		qu.treeDepths[pRoot]++
	}
	qu.count--
}

// find the root of p
func (qu *heightQuickUnionUF) Find(p int) int {
	if l := len(qu.roots); p >= l {
		return -1
	}
	for p != qu.roots[p] {
		p = qu.roots[p]
	}
	return p
}

// if contact p and contact q in the same weight, return true
func (qu *heightQuickUnionUF) Connected(p, q int) bool {
	if l := len(qu.roots); p >= l || q >= l {
		return false
	}
	return qu.Find(p) == qu.Find(q)
}

func (qu *heightQuickUnionUF) Count() int {
	return qu.count
}
