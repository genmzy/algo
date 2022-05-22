### sample data

- sample data result:
```log
goos: linux
goarch: amd64
pkg: algo/practice/ch1/1_4/1_4_44_45
cpu: AMD Ryzen 7 4800H with Radeon Graphics
BenchmarkDupBirth-16          590982          1817 ns/op             831 B/op          4 allocs/op
--- BENCH: BenchmarkDupBirth-16
	randdup_test.go:12: average 23 people create duplicate birthday
	randdup_test.go:12: average 22 people create duplicate birthday
	randdup_test.go:12: average 24 people create duplicate birthday
	randdup_test.go:12: average 23 people create duplicate birthday
BenchmarkCollectAll-16           174       7120888 ns/op          846461 B/op        574 allocs/op
--- BENCH: BenchmarkCollectAll-16
	randdup_test.go:21: average 171904 times collect all numbers of 20000
	randdup_test.go:21: average 184667 times collect all numbers of 20000
	randdup_test.go:21: average 188649 times collect all numbers of 20000
PASS
ok      algo/practice/ch1/1_4/1_4_44_45 3.038s
```
