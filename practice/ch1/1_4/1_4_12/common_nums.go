package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func CommonNums(a, b []int) []int {
	maxLen := max(len(a), len(b))
	res := make([]int, 0, maxLen)
	j := 0
	for i := 0; i < len(a); i++ {
		for j < len(b) && a[i] > b[j] {
			j++
		}
		if j == len(b) {
			break
		}
		if a[i] == b[j] {
			res = append(res, a[i])
		}
	}
	return res
}
