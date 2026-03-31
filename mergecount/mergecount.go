package main

import (
	"fmt"
	"math"
)

func mergesort(arr []int, low, high int) int {
	if low < high {
		mid := (low + high) / 2
		leftCount := mergesort(arr, low, mid)
		rightCount := mergesort(arr, mid+1, high)
		mergeCount := merge(arr, low, mid, high)
		return leftCount + rightCount + mergeCount
	}
	return 0
}

func merge(arr []int, low, mid, high int) int {
	left := append([]int(nil), arr[low:mid+1]...)
	right := append([]int(nil), arr[mid+1:high+1]...)

	left = append(left, math.MaxInt64)
	right = append(right, math.MaxInt64)

	i, j := 0, 0
	count := 0
	leftLen := mid - low + 1
	for k := low; k <= high; k++ {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
			count += leftLen - i
		}
	}
	return count
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	count := mergesort(arr, 0, n-1)
	fmt.Println(count)

}
