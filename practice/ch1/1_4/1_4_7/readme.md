**Answer**

```go
func ThreeSumCnt(nums []int) int {
	res := 0
	n := len(num)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res++
				}
			}
		}
	}
	return res
}
```

| express                 | times |
|-------------------------|-------|
| i<n                     | n     |
| j<n                     | n²    |
| k<n                     | n³    |
| res++                   | n³    |
| nums[i]+nums[j]+nums[k] | n³    |
> so, n³
