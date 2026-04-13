package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0
	for left < right {
		h := min(height[left], height[right])
		water := h * (right - left)
		if water > maxWater {
			maxWater = water
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxWater := maxArea(height)
	fmt.Println(maxWater)

	height = []int{1, 1}
	maxWater = maxArea(height)
	fmt.Println(maxWater)
}
