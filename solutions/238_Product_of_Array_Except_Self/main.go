package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	result := make([]int, n)

	leftProduct := 1
	for i := 0; i < n; i++ {
		result[i] = leftProduct
		leftProduct *= nums[i]
	}

	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return result
}

func main() {
	arr := []int{1, 2, 3, 4}
	result := productExceptSelf(arr)
	fmt.Println(result) // Output: [24, 12, 8, 6]

	arr = []int{2, 3, 4, 5}
	result = productExceptSelf(arr)
	fmt.Println(result) // Output: [60, 40, 30, 24]
}
