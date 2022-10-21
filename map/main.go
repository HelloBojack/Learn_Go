package main

import "fmt"

func main() {
	m := make(map[string]string, 3)
	m["one"] = "js"
	m2 := map[string]string{
		"two":   "js",
		"three": "java",
	}
	fmt.Printf("%v %v\n", m, m2)

	for key, value := range m2 {
		fmt.Printf("%v %v\n", key, value)
	}
}
