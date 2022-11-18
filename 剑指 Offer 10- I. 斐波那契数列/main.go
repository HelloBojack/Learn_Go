package main

import "fmt"

func fib(n int) int {
	const mod int = 1e9 + 7
	if n < 2 {
		return n
	}
	q, p, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		q = p
		p = r
		r = (q + p) % mod
	}

	return r
}

func main() {
	dp := fib(10)
	fmt.Println(dp)
}
