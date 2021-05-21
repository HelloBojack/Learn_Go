package main

import "fmt"

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	// 1 1 2 3 5
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fib(2))
}
