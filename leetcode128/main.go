package main

import "fmt"

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	longest := 0
	for num := range set {
		if !set[num-1] {
			current := num
			length := 1
			for set[current+1] {
				current++
				length++
			}
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
	nums = []int{1, 0, 1, 2}
	fmt.Println(longestConsecutive(nums))
}
