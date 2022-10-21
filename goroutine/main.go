package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine over...")

		fmt.Println("goroutine ing...")

		c <- 123
	}()

	num := <-c

	fmt.Println(num)

	fmt.Println("main over...")

}
