package main

import "fmt"

// return 先于 defer
func fn(i int) int {
	defer func() {
		i++
		fmt.Println(i)
	}()

	return i
}

func main() {
	i := fn(1)
	fmt.Println(i)
}
