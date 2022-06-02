package unionfind

type UnionFind interface {
	Union(p, q int)
	Find(p int) int
	Connected(p, q int) bool
	Count() int
}

type UnionFindWithVisit interface {
	UnionFind
	FindWithVisit(p int) (int, int)
	UnionWithVisit(p, q int) int
}
