package main

import "fmt"

func main() {
	s := make([]string, 3)
	s[0] = "0"
	s = append(s, "1")

	s2 := []string{"1", "2", "3"}
	s3 := make([]string, len(s2))
	copy(s3, s2)
	twoDs := make([][]string, 3)
	twoDs[0] = s
	twoDs[1] = s2
	twoDs[2] = s3
	fmt.Println(s, s2, s3, twoDs)
}
