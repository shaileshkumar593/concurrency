package main

import "fmt"

// nextGreaterSteps returns steps to next greater element for each index
func nextGreaterSteps(arr []int) []int {
	n := len(arr)
	res := make([]int, n)

	stack := []int{} // store indices

	// Traverse from right to left
	for i := n - 1; i >= 0; i-- {

		// Remove all elements <= current
		for len(stack) > 0 && arr[stack[len(stack)-1]] <= arr[i] {
			stack = stack[:len(stack)-1]
		}

		// Assign result
		if len(stack) == 0 {
			res[i] = 0
		} else {
			res[i] = stack[len(stack)-1] - i
		}

		// Push current index
		stack = append(stack, i)
	}

	return res
}

func main() {
	arr := []int{84, 74, 21, 19, 75, 66, 44, 55, 70, 34, 32, 29, 92, 4, 6, 105, 101}

	result := nextGreaterSteps(arr)

	fmt.Println("Input :", arr)
	fmt.Println("Output:", result)
}
