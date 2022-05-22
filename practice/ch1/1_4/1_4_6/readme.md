**Answer**

---

- a
```go
sum := 0
for n := N; n > 0; n /= 2 {
	for i := 0; i < n; i++ {
		sum++
	}
}
```
	> = N+N/2+N/4+N/8+ ... = N(1+1/2+1/4 + ... ) ~ N(1+ 1-1/2 + 1/2-1/4 + 1/4-1/8 + 1/8-1/16 + ...) ~ 2N

- b
```go
sum := 0
for i := 1; i < N; i *= 2 {
	for j := 0; j < i; j++ {
		sum++
	}
}
```
	> ~ 2N (same above)

- c
```go
sum := 0
for i := 1; i < N; i *= 2 {
	for j := 0; j < N; j++ {
		sum++
	}
}
```
	> N*lgN
