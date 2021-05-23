package main

import "fmt"

func swap(aa *int, bb *int) {
	temp := *aa
	*aa = *bb
	*bb = temp
}

func main() {
	a, b := 10, 20

	a, b = b, a
	fmt.Println(a, b)

	swap(&a, &b)
	fmt.Println(a, b)
}
