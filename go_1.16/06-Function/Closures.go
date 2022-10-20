package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextSeq := intSeq()

	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
}
