package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	k, v := m["one"]

	m2 := map[string]int{"three": 3, "four": 4}

	fmt.Println(m, k, v, m2)
}
