**Question**
- Give a formula to predict the running time of a program for a problem of size N when doubling experiments have shown that the doubling factor is 2^b and the running time for problems of size N0 is T.

**Answer**

- *T[N]* = 2^b * *T[N/2]*;
- *T[N0]* = T

> T[N] = 2^b * T[N/2] = 2^(N/(N/2)*b) * T[N/4] = 2^(2*b)*T[N/4]
> = 2^( 2 * ((N/2)/(N/4)) * b) * T[N/8]
> = 2^( 2 * 2 * b) * T[N/8]
> ...
> = 2^(b*2^x) * T[N/2^x]
>
> let N/2^x = N0 => 2^x=N/N0 => x=lg(N/N0)
>
> so: T[N] = 2^(2^lg(N/N0))*T[N0] = T*2^(N/N0)
