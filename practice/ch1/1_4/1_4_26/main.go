package main

type Point struct {
	x, y int
}

// 1.4.26 3-collinearity.
//
// Suppose that you have an algorithm that takes as input N
// distinct points in the plane and can return the number
// of triples that fall on the same line.
// Show that you can use this algorithm to solve the 3-sum problem.
//
// Strong hint: Use algebra to show that (a, a^3), (b, b^3),
// and (c, c^3) are collinear if and only if a + b + c = 0.
//
// return the three point collinears of all points
func FindThreePointsCollinear(points []Point) int {
	// TODO: I don't figure this out <17-05-22, genmzy> //
	return -1
}
