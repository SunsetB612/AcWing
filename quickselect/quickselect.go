package main

import (
	"fmt"
)

func quickselect(arr []int, k, low, high int) int {
	if low == high {
		return arr[low]
	}
	p := partition(arr, low, high)
	if k <= p {
		return quickselect(arr, k, low, p)
	} else {
		return quickselect(arr, k, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[(low+high)/2]
	i, j := low-1, high+1
	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}
		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			return j
		}
	}
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	target := quickselect(arr, k-1, 0, n-1)

	fmt.Println(target)
}
