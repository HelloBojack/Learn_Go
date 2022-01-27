package main

import "fmt"

// slice 切片 动态数组 可以动态长度 函数传递为引用传递 可直接修改原数组
func sliceFn(s []string) {
	s[2] = "200"
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	s := make([]string, 3)
	s[0] = "0"
	s = append(s, "1")

	s2 := []string{"1", "2", "3"}
	sliceFn(s2)

	s3 := make([]string, len(s2))
	copy(s3, s2)

	twoDs := make([][]string, 3)
	twoDs[0] = s
	twoDs[1] = s2
	twoDs[2] = s3
	fmt.Println(s, s2, s3, twoDs)
}
