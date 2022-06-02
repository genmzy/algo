package unionfind

func weightQuickUnion(n int) *weightQuickUnionUF {
	qu := &weightQuickUnionUF{
		roots: make([]int, 0, n),
		count: n,
	}
	// why initialize components like this?
	//    to store contacts components as all contacts are not
	//    in one component
	for i := 0; i < qu.count; i++ {
		qu.roots = append(qu.roots, i)
	}
	qu.treeSizes = make([]int, n, n)
	for i := 0; i < n; i++ {
		qu.treeSizes[i] = 1
	}
	return qu
}

// improved union find algorithm: make the
func WeightQuickUnionUF(n int) UnionFind {
	return weightQuickUnion(n)
}

func WeightQuickUnionUFWithVisit(n int) UnionFindWithVisit {
	return weightQuickUnion(n)
}

type weightQuickUnionUF struct {
	roots     []int // contact id => root contact id
	treeSizes []int // root contact id => size of tree of root contact id
	count     int   // component count
}

func (qu *weightQuickUnionUF) Union(p, q int) {
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
