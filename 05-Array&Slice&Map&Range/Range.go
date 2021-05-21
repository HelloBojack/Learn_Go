package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, i := range a {
		sum += i
	}
	fmt.Println(sum)

	m := map[string]string{"one": "1", "two": "2", "three": "3"}
	for k, v := range m {
		fmt.Printf("%s->%s\n", k, v)
	}
}
