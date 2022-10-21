package main

import "fmt"

func showSlice(s1 []int) []int {
	return s1
}

func main() {
	s2 := showSlice([]int{1, 2, 3})
	fmt.Println(s2)

	s3 := make([]int, 3, 5)

	fmt.Printf("%d %d %v\n", len(s3), cap(s3), s3)

	s3 = append(s3, 1)
	s3 = append(s3, 2)
	s3 = append(s3, 3)
	s3 = append(s3, 4)

	fmt.Printf("%d %d %v\n", len(s3), cap(s3), s3)

	s4 := s3[0:2]
	s5 := make([]int, 3)
	s4[0] = 4

	copy(s5, s4)
	fmt.Printf("%v %v %v\n", s4, s3, s5)

}
