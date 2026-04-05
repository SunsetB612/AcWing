package main

import "fmt"

//func myPow(x float64, n int) float64 {
//	if n == 0 {
//		return 1
//	}
//	if n < 0 {
//		return myPow(1/x, -n)
//	}
//	if n%2 == 1 {
//		return x * myPow(x, n-1)
//	}
//	half := myPow(x, n/2)
//	return half * half
//}

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	result := 1.0
	for n > 0 {
		if n&1 == 1 {
			result *= x
		}
		x *= x
		n >>= 1
	}
	return result
}

func main() {
	fmt.Printf("%.5f\n", myPow(2.00000, 10))
	fmt.Printf("%.5f\n", myPow(2.10000, 3))
	fmt.Printf("%.5f\n", myPow(2.00000, -2))
}
