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

func all2(a, b, c int) (r1 int, r2 int) {
	r1 = 1
	r2 = 2
	return
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
	fmt.Println(all2(1, 2, 3))
}
