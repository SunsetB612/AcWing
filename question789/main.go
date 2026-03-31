package main

import "fmt"

func lowerBound(arr []int, target int) int {
	low, high := 0, len(arr)-1
	result := -1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			result = mid
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}

func upperBound(arr []int, target int) int {
	low, high := 0, len(arr)-1
	result := -1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			result = mid
			low = mid + 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result
}
func main() {
	var n, q int
	fmt.Scan(&n, &q)

	arr := make([]int, n)
	target := make([]int, q)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < q; i++ {
		fmt.Scan(&target[i])
	}

	for _, v := range target {
		fmt.Println(lowerBound(arr, v), upperBound(arr, v))
	}

}
