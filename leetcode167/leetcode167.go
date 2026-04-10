package main

import "fmt"

func twoSum(numbers []int, target int) []int {

	n := len(numbers)
	i, j := 0, n-1
	ans := []int{-1, -1}
	for i < n && j >= 0 {
		sum := numbers[i] + numbers[j]
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			ans[0], ans[1] = i+1, j+1
			break
		}
	}
	return ans
}

func main() {
	numbers := []int{-1, 0}
	target := -1
	fmt.Println(twoSum(numbers, target))
}
