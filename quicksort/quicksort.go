package main

import "fmt"

func quicksort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quicksort(arr, low, pivot)
		quicksort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := (arr[low] + arr[high]) / 2
	i := low - 1
	j := high + 1
	for {
		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}
		for {
			i++
			if arr[i] >= pivot {
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
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	quicksort(arr, 0, len(arr)-1)

	for _, v := range arr {
		fmt.Print(v, " ")
	}
}
