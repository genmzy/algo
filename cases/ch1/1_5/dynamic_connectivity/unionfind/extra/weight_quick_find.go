package uf_extra

import "algo/cases/ch1/1_5/dynamic_connectivity/unionfind"

func WeightQuickFindUF(n int) unionfind.UnionFind {
	qf := &weightQuickFindUF{
		components: make([]int, 0, n),
		count:      n,
	}
	for i := 0; i < qf.count; i++ {
		qf.components = append(qf.components, i)
	}
	for i := 0; i < qf.count; i++ {
		qf.sizes[i] = 1
	}
	return qf
}

type weightQuickFindUF struct {
	components []int
	sizes      []int
	count      int
}

func (qf *weightQuickFindUF) Union(p, q int) {
	if l := len(qf.components); p >= l || q >= l {
		return
	}
	pId := qf.Find(p)
	qId := qf.Find(q)
	if pId == qId {
		return
	}
	if qf.sizes[pId] < qf.sizes[qId] {
		for i := 0; i < len(qf.components); i++ {
			if qf.components[i] == pId {
				qf.components[i] = qId
			}
		}
		qf.sizes[qId] += qf.sizes[pId]
	} else {
		for i := 0; i < len(qf.components); i++ {
			if qf.components[i] == qId {
				qf.components[i] = pId
			}
		}
		qf.sizes[pId] += qf.sizes[qId]
	}
	qf.count--
}

func (qf *weightQuickFindUF) Find(p int) int {
	if l := len(qf.components); p >= l {
		return -1
	}
	return qf.components[p]
}

// if contact p and contact q in the same weight, return true
func (qf *weightQuickFindUF) Connected(p, q int) bool {
	if l := len(qf.components); p >= l || q >= l {
		return false
	}
	return qf.Find(p) == qf.Find(q)
}

func (qf *weightQuickFindUF) Count() int {
	return qf.count
}
