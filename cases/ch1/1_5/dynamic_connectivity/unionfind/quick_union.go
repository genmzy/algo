package unionfind

func quickUnion(n int) *quickUnionUF {
	qu := &quickUnionUF{
		roots: make([]int, 0, n),
		count: n,
	}
	// why initialize components like this?
	//    to store contacts components as all contacts are not
	//    in one component
	for i := 0; i < qu.count; i++ {
		qu.roots = append(qu.roots, i)
	}
	return qu
}

func QuickUnionUF(n int) UnionFind {
	return quickUnion(n)
}

func QuickUnionUFWithVisit(n int) UnionFindWithVisit {
	return quickUnion(n)
}

type quickUnionUF struct {
	roots []int // contact id => root contact id
	count int   // component count
}

// NOTE: here is a question, why not p connect q instead of root of p
// connect the root of q?
// if just connect p and q, the roots will not be a tree, it's a graph
// it will hard to judge if two point is connected
//
// the fact is what we do here is make the graph be a tree to
// find the path easily
func (qu *quickUnionUF) Union(p, q int) {
	if l := len(qu.roots); p >= l || q >= l {
		return
	}
	// find the root of p and q
	pRoot := qu.Find(p)
	qRoot := qu.Find(q)
	if pRoot == qRoot {
		return
	}
	// set q's root as q's root's root(original is itself),
	// which means merge the tree of p's root into
	// the tree of q's root
	//
	//                (qRoot)             (qRoot)
	//                   *                *
	//        (pRoot)  /  \   =>        / | \
	//          *     *   *     (pRoot)*  * *
	//        / |                    /  \
	//       *  *                   *   *
	qu.roots[pRoot] = qRoot
	qu.count--
}

func (qu *quickUnionUF) UnionWithVisit(p, q int) int {
	if l := len(qu.roots); p >= l || q >= l {
		return 0
	}
	v := 0 // visit times
	// find the root of p and q
	pRoot, k := qu.FindWithVisit(p)
	v += k
	qRoot, k := qu.FindWithVisit(q)
	v += k
	if pRoot == qRoot {
		return v
	}
	qu.roots[pRoot] = qRoot
	v++
	qu.count--
	return v
}

// find the root of p
func (qu *quickUnionUF) Find(p int) int {
	if l := len(qu.roots); p >= l {
		return -1
	}
	for p != qu.roots[p] {
		p = qu.roots[p]
	}
	return p
}

func (qu *quickUnionUF) FindWithVisit(p int) (int, int) {
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
func (qu *quickUnionUF) Connected(p, q int) bool {
	if l := len(qu.roots); p >= l || q >= l {
		return false
	}
	return qu.Find(p) == qu.Find(q)
}

func (qu *quickUnionUF) Count() int {
	return qu.count
}
