package main

func nextGreaterElements(nums []int) []int {
	leng := len(nums)
	res := make([]int, leng)

	for i := range res {
		res[i] = -1
	}

	stack := []int{}

	for i := 0; i < leng*2-1; i++ {
		num := nums[i%leng]

		for len(stack) > 0 && num > nums[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = num

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i%leng)
	}

	return res
}

func main() {

}
