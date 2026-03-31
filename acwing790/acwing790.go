package main

import (
	"fmt"
)

func cubeRoot(n float64) float64 {
	low, high := -10000.00, 10000.00
	for high-low > 1e-8 {
		mid := (low + high) / 2
		if mid*mid*mid < n {
			low = mid
		} else {
			high = mid
		}
	}
	return low
}

func main() {
	var n float64
	fmt.Scan(&n)
	fmt.Printf("%.6f", cubeRoot(n))
}
