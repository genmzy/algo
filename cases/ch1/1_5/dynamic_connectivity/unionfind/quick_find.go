// algorithm 1.5
package unionfind

func quickFind(n int) *quickFindUF {
	qf := &quickFindUF{
		components: make([]int, 0, n),
		count:      n,
	}
	// why initialize components like this?
	//    to store contacts components as all contacts are not
	//    in one component
	for i := 0; i < qf.count; i++ {
		qf.components = append(qf.components, i)
	}
	return qf
}

func QuickFindUF(n int) UnionFind {
	return quickFind(n)
}

func QuickFindUFWithVisit(n int) UnionFindWithVisit {
	return quickFind(n)
}

type quickFindUF struct {
	components []int // contact id => component id
	count      int   // component count
}

// connect contact p and contact q
func (qf *quickFindUF) Union(p, q int) {
	if l := len(qf.components); p >= l || q >= l {
		return
	}
	// got id of components of p and q
	pId := qf.Find(p)
	qId := qf.Find(q)
	if pId == qId {
		return
	}
	for i := 0; i < len(qf.components); i++ {
		if qf.components[i] == pId {
			qf.components[i] = qId
		}
	}
	qf.count--
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

// return the component id of p
func (qf *quickFindUF) Find(p int) int {
	if l := len(qf.components); p >= l {
		return -1
	}
	return qf.components[p]
}

func (qf *quickFindUF) FindWithVisit(p int) (int, int) {
	if l := len(qf.components); p >= l {
		return -1, 0
	}
	return qf.components[p], 1
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
