package main

import "fmt"

// 数组作为参数 必须明确数组长度
func arrFn(arr [5]int) {
	// 且传参是值传递，不修改原数组
	arr[4] = 1
	for index, value := range arr {
		fmt.Println(index, value)
	}
}

func main() {
	// 数组：固定长度
	var arr [5]int
	arr1 := [5]int{1, 2, 3}

	arrFn(arr)
	arrFn(arr1)
	fmt.Println(arr)
}
