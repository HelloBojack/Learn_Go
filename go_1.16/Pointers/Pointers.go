package main

import "fmt"

func changeValue(number int) {
	number = 123
}
func changePointer(number *int) {
	*number = 123
}
func main() {
	i := 1
	fmt.Println(&i)
	changeValue(i)
	fmt.Println(i)
	changePointer(&i)
	fmt.Println(i)
}
