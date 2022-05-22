- code
```go
// O(N*lgN)
func twoSumFast(nums []int) [][2]int {
	sort.Ints(nums) // <- A
	res := make([][2]int, 0, 300000) // <- B
	n := len(nums)
	for i := 0; i < n; i++ { // <- C
		if nums[i] > 0 {
			break
		}
		if j := binarysearch.Rank(-nums[i], nums); j > i { // <-D
			res = append(res, [2]int{nums[i], nums[j]})
		}
	}
	return res
}
```
| Block       | Runtime   | frequence | total time                |
| ----------- | --------- | --------- | ------                    |
| D           | t0        | lgN       | t0*lgN                    |
| C           | t1        | N/2       | t1*N/2                    |
| B           | t2        | 1         | t2                        |
| A           | t3        | N*lgN     | t3*N*lgN                  |
|             |           |           |                           |
|             |           |           | (t3+t1)*N*lgN+(t1/2)*N+t2 |
|             |           |           | ~(t3+t1)/N*lgN            |
|             |           |           | N*lgN                     |
