package main

import "fmt"

func isHappy(n int) bool {
	slow, fast := n, next(n)
	for fast != 1 && slow != fast {
		slow = next(slow)
		fast = next(next(fast))
	}
	return fast == 1
}

func next(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n = n / 10
	}
	return sum
}

func main() {
	n := 2
	fmt.Println(isHappy(n))
}
