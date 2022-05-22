### for golang, compare interface{} and T any and int performance cost

- result:
	```log
	goos: linux
	goarch: amd64
	pkg: algo/practice/ch1/1_4/1_4_37
	cpu: AMD Ryzen 7 4800H with Radeon Graphics
	BenchmarkInt-16         77590006            13.97 ns/op        8 B/op          0 allocs/op
	BenchmarkGeneric-16     84303668            13.94 ns/op        8 B/op          0 allocs/op
	BenchmarkIface-16       46113264            24.79 ns/op       16 B/op          0 allocs/op
	PASS
	ok      algo/practice/ch1/1_4/1_4_37    3.474s

	```
