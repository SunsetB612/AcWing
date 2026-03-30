package main

import (
	"fmt"
	"math"
)

func mergesort(arr []int, low, high int) {
	if low < high {
		mid := (low + high) / 2
		mergesort(arr, low, mid)
		mergesort(arr, mid+1, high)
		merge(arr, low, mid, high)
	}
}

func merge(arr []int, low, mid, high int) {
	left := append([]int(nil), arr[low:mid+1]...)
	right := append([]int(nil), arr[mid+1:high+1]...)

	left = append(left, math.MaxInt64)
	right = append(right, math.MaxInt64)

	i, j := 0, 0
	for k := low; k <= high; k++ {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	mergesort(arr, 0, len(arr)-1)

	for _, v := range arr {
		fmt.Print(v, " ")
	}
}
