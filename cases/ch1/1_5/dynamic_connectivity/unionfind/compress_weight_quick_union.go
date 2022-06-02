package unionfind

// improved union find algorithm: make the
func compressWeightQuickUnion(n int) *compressWeightQuickUnionUF {
	qu := &compressWeightQuickUnionUF{
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

func CompressWeightQuickUnionUF(n int) UnionFind {
	return compressWeightQuickUnion(n)
}

func CompressWeightQuickUnionUFWithVisit(n int) UnionFindWithVisit {
	return compressWeightQuickUnion(n)
}

type compressWeightQuickUnionUF struct {
	roots     []int // contact id => root contact id
	treeSizes []int // root contact id => size of tree of root contact id
	count     int   // component count
}

func (qu *compressWeightQuickUnionUF) Union(p, q int) {
	if l := len(qu.roots); p >= l || q >= l {
		return
	}
	pRoot := qu.find(p)
	qRoot := qu.find(q)
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

func (qu *compressWeightQuickUnionUF) UnionWithVisit(p, q int) int {
	if l := len(qu.roots); p >= l || q >= l {
		return 0
	}
	v := 0
	//pRoot := qu.find(p)
	//qRoot := qu.find(q)
	pRoot, k := qu.findWithVisit(p)
	v += k
	qRoot, k := qu.findWithVisit(q)
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

func (qu *compressWeightQuickUnionUF) find(p int) int {
	k := qu.Find(p)
	for p != qu.roots[p] {
		pPrev := p
		p = qu.roots[p]
		qu.roots[pPrev] = k
	}
	return k
}

func (qu *compressWeightQuickUnionUF) findWithVisit(p int) (int, int) {
	k, v := qu.FindWithVisit(p)
	for p != qu.roots[p] {
		pPrev := p
		p = qu.roots[p]
		qu.roots[pPrev] = k
		v += 3
	}
	v += 1
	return k, v
}

// find the root of p
func (qu *compressWeightQuickUnionUF) Find(p int) int {
	if l := len(qu.roots); p >= l {
		return -1
	}
	for p != qu.roots[p] {
		p = qu.roots[p]
	}
	return p
}

func (qu *compressWeightQuickUnionUF) FindWithVisit(p int) (int, int) {
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

// if contact p and contact q in the same pathCompressWeight, return true
func (qu *compressWeightQuickUnionUF) Connected(p, q int) bool {
	if l := len(qu.roots); p >= l || q >= l {
		return false
	}
	return qu.Find(p) == qu.Find(q)
}

func (qu *compressWeightQuickUnionUF) Count() int {
	return qu.count
}
