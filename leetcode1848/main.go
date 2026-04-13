package main

import (
	"fmt"
	"math"
)

func getMinDistance(nums []int, target int, start int) int {
	var indexs []int
	for i, num := range nums {
		if num == target {
			indexs = append(indexs, i)
		}
	}
	answer := math.MaxInt64
	for _, index := range indexs {
		result := int(math.Abs(float64(index - start)))
		if result < answer {
			answer = result
		}
	}
	return answer
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 5
	start := 3
	fmt.Println(getMinDistance(nums, target, start))

	nums = []int{1}
	target = 1
	start = 0
	fmt.Println(getMinDistance(nums, target, start))

	nums = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	target = 1
	start = 0
	fmt.Println(getMinDistance(nums, target, start))
}
