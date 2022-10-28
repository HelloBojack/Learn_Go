package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := []int{}
	res := make([]int, n)

	for i := 0; i < n; i++ {
		num := temperatures[i]

		for len(stack) > 0 && num > temperatures[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)

	}
	return res
}

func main() {

}
