package main

import "fmt"

func binarysearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13}
	target := 13

	index := binarysearch(arr, target)

	if index != -1 {
		fmt.Println(index)
	} else {
		fmt.Println("Not found")
	}
}
