package localminimum

// 1.4.18 Local minimum of an array.
// Write a program that, given an array a[] of N distinct integers, finds a local minimum: an index i such that
// a[i-1] < a[i] < a[i+1].
// Your program should use ~2 lg N compares in the worst case.
func ArrayLocalMinimum(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 || nums[0] < nums[1] {
		return 0
	}
	l, r := 1, n-2
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] <= nums[mid-1] && nums[mid] <= nums[mid+1] {
			return mid
		} else {
			if nums[mid-1] > nums[mid+1] {
				l = mid
			} else {
				r = mid
			}
		}
	}
	return -1
}

/*
	1.4.19 Local minimum of a matrix.
	Given an N-by-N array a[] of N^2 distinct integers, design an algorithm that runs in time proportional
	to N to find a local minimum: a pair of indices i and j such that
		a[i][j] < a[i+1][j],
		a[i][j] < a[i][j+1],
		a[i][j] < a[i-1][j],
	and a[i][j] < a[i][j-1].
	The running time of your program should be proportional to N in the worst case.
	length of array must equals
*/

// basic thought: find a local minimum in every array of matrix
// and put them into an array `b', then find the local minimum in `b'
func MatrixLocalMimium(matrix [][]int) (int, int) {
	n := len(matrix) // assert len(matrix[i]) == n
	if n == 0 {
		return -1, -1
	}
	if n == 1 || (matrix[0][0] < matrix[1][0] &&
		matrix[0][0] < matrix[0][1] &&
		matrix[0][0] < matrix[1][1]) {
		return 0, 0
	}
	colmax := make([]int, 0, n)
	colmaxIdx := make([]int, 0, n)
	for _, v := range matrix { // <- O(N)
		j := ArrayLocalMinimum(v) // <- O(lgN)
		if j == -1 {
			continue
		}
		colmaxIdx = append(colmaxIdx, j)
		colmax = append(colmax, v[j])
	}
	i := ArrayLocalMinimum(colmax) // <- O(lgN)
	if i == -1 {
		return -1, -1
	}
	return i, colmaxIdx[i]
}

type pos struct {
	x, y int
}

func (p pos) v(matrix [][]int) int {
	return matrix[p.x][p.y]
}

func (p pos) down() pos {
	return pos{p.x, p.y - 1}
}

func (p pos) left() pos {
	return pos{p.x - 1, p.y}
}

func (p pos) up() pos {
	return pos{p.x, p.y + 1}
}

func (p pos) right() pos {
	return pos{p.x + 1, p.y}
}

func posLessThan(a, b pos) bool {
	if a.x < b.x && a.y < b.y {
		return true
	}
	return false
}

func mid(a, d pos) pos {
	midX := a.x + (d.x-a.x)>>1
	midY := a.y + (d.y-a.y)>>1
	return pos{midX, midY}
}

// Divide & Conquer
// Look at boundary, center row, and center column
// Find global minimum within
// can't be in window
// Recursive in quadrant, including green boundary
//
// Lemma: if you enter a quadrant, it contains a peak of the overall array
// Invariant: minimum element of window never increase as we scend in recursion
// Theorem: valley in visited quadrant is also valley in overall array
func MatrixLocalMinimumFaster(matrix [][]int) (int, int) {
	n := len(matrix)
	if n == 0 {
		return -1, -1
	}
	if n == 1 || (matrix[0][0] < matrix[1][0] &&
		matrix[0][0] < matrix[0][1] &&
		matrix[0][0] < matrix[1][1]) {
		return 0, 0
	}
	return matrixLocalMinimum(matrix, pos{0, 0}, pos{n - 1, n - 1})
}

type direction int

const (
	current direction = iota // itself
	up
	down
	left
	right
)

func (p pos) isLocalMinimum(matrix [][]int) bool {
	n := len(matrix) - 1
	left, right, up, down := false, false, false, false
	if p.x > 0 {
		left = matrix[p.x][p.y] <= matrix[p.x-1][p.y]
	} else {
		left = true
	}
	if p.x < n {
		right = matrix[p.x][p.y] <= matrix[p.x+1][p.y]
	} else {
		right = true
	}
	if p.y < n {
		up = matrix[p.x][p.y] <= matrix[p.x][p.y+1]
	} else {
		up = true
	}
	if p.y > 0 {
		down = matrix[p.x][p.y] <= matrix[p.x][p.y-1]
	} else {
		down = true
	}
	return left && right && up && down
}

// TODO: fix boundary condition <2022-05-13, genmzy> //
func matrixLocalMinimum(matrix [][]int, m, n pos) (int, int) {
	if m == n {
		return -1, -1
	}
	p := mid(m, n)
	if p == m {
		for _, v := range []pos{n, m, {m.x, n.y}, {n.x, m.y}} {
			if v.isLocalMinimum(matrix) {
				return v.x, v.y
			}
		}
		return -1, -1
	}
	if p.v(matrix) == p.up().v(matrix) &&
		p.v(matrix) == p.down().v(matrix) &&
		p.v(matrix) == p.left().v(matrix) &&
		p.v(matrix) == p.right().v(matrix) {
		return p.x, p.y
	}
	curMinimum, column, line := false, up, left
	columnPos, linePos := p.up(), p.left()
	if p.up().v(matrix) > p.down().v(matrix) {
		column = down
		columnPos = p.down()
	}
	if p.left().v(matrix) > p.right().v(matrix) {
		line = right
		columnPos = p.right()
	}
	if p.v(matrix) <= columnPos.v(matrix) && p.v(matrix) <= linePos.v(matrix) {
		curMinimum = true
	}
	if curMinimum {
		return p.x, p.y
	}
	if column == up {
		if line == right {
			return matrixLocalMinimum(matrix, p, n)
		} else {
			return matrixLocalMinimum(matrix, pos{m.x, p.y}, pos{p.x, n.y}) // left up
		}
	} else {
		if line == right {
			return matrixLocalMinimum(matrix, pos{p.x, p.y}, pos{n.x, p.y})
		} else {
			return matrixLocalMinimum(matrix, m, p)
		}
	}
}
