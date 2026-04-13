package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	for _, num1 := range nums1 {
		set[num1] = true
	}

	resultSet := make(map[int]bool)
	for _, num2 := range nums2 {
		if set[num2] {
			resultSet[num2] = true
		}
	}

	var result []int
	for k := range resultSet {
		result = append(result, k)
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}
