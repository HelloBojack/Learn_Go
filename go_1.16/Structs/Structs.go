package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	bojack := Person{"Bojack", 24}
	fmt.Println(bojack)
	bojack2 := &bojack
	fmt.Println(bojack2.name)
}
