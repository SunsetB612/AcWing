package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func intersect(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]int)
	set2 := make(map[int]int)

	for _, num1 := range nums1 {
		set1[num1]++
	}

	for _, num2 := range nums2 {
		set2[num2]++
	}

	resultSet := make(map[int]int)
	for num := range set1 {
		if _, ok := set2[num]; ok {
			resultSet[num] = min(set1[num], set2[num])
		}
	}

	var result []int
	for k, v := range resultSet {
		for i := 0; i < v; i++ {
			result = append(result, k)
		}
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersect(nums1, nums2))

	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}
