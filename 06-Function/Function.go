package main

import "fmt"

func plus(a, b, c int) int {
	return a + b + c
}
func sub(a, b, c int) int {
	return a - b - c
}

func allInOne(a, b, c int) (int, int) {
	return plus(a, b, c), sub(a, b, c)
}

func plusPlus(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(allInOne(1, 2, 3))
	fmt.Println(plusPlus(1, 2, 3, 4, 5, 6, 7, 8))
}
