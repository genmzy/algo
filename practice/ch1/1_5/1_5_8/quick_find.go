package main

import "algo/cases/ch1/1_5/dynamic_connectivity/unionfind"

func QuickFindUF(n int) unionfind.UnionFind {
	qf := &quickFindUF{
		components: make([]int, 0, n),
		count:      n,
	}
	for i := 0; i < qf.count; i++ {
		qf.components = append(qf.components, i)
	}
	return qf
}

type quickFindUF struct {
	components []int
	count      int
}

// XXX: wrong union implementation
func (qf *quickFindUF) Union(p, q int) {
	if l := len(qf.components); p >= l || q >= l {
		return
	}
	if qf.Connected(p, q) {
		return
	}
	for i := 0; i < len(qf.components); i++ {
		if qf.components[i] == qf.components[p] {
			qf.components[i] = qf.components[q]
		}
	}
}

// return the component id of p
func (qf *quickFindUF) Find(p int) int {
	if l := len(qf.components); p >= l {
		return -1
	}
	return qf.components[p]
}

// if contact p and contact q in the same weight, return true
func (qf *quickFindUF) Connected(p, q int) bool {
	if l := len(qf.components); p >= l || q >= l {
		return false
	}
	return qf.Find(p) == qf.Find(q)
}

func (qf *quickFindUF) Count() int {
	return qf.count
}

func (qf *quickFindUF) UnionWithVisit(p, q int) int {
	if l := len(qf.components); p >= l || q >= l {
		return 0
	}
	v := 0
	// got id of components of p and q
	pId, k := qf.FindWithVisit(p)
	v += k
	qId, k := qf.FindWithVisit(q)
	v += k
	if pId == qId {
		return v
	}
	for i := 0; i < len(qf.components); i++ {
		if qf.components[i] == pId {
			qf.components[i] = qId
			v += 1
		}
		v += 1
	}
	qf.count--
	return v
}

func (qf *quickFindUF) FindWithVisit(p int) (int, int) {
	if l := len(qf.components); p >= l {
		return -1, 0
	}
	return qf.components[p], 1
}
