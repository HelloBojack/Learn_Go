package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := map[int]int{}
	stack := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			fmt.Println(stack, len(stack)-1)
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			mp[num] = stack[len(stack)-1]
		} else {
			mp[num] = -1
		}
		stack = append(stack, num)
	}

	res := []int{}
	for _, value := range nums1 {
		if value, ok := mp[value]; ok {
			res = append(res, value)
		}
	}

	return res
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	res := nextGreaterElement(nums1, nums2)
	fmt.Println(res)
}
