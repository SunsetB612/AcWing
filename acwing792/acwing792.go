package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	negetive := false
	if len(a) < len(b) || (len(a) == len(b) && a < b) {
		a, b = b, a
		negetive = true
	}

	for len(b) < len(a) {
		b = "0" + b
	}

	n := len(a)
	borrow := 0
	result := make([]byte, n)
	for i := n - 1; i >= 0; i-- {
		diff := int(a[i]-'0') - int(b[i]-'0') - borrow
		if diff < 0 {
			borrow = 1
			diff += 10
		} else {
			borrow = 0
		}
		result[i] = byte(diff) + '0'
	}

	start := 0
	for start < n-1 && result[start] == '0' {
		start++
	}

	if negetive {
		fmt.Printf("-")
	}
	fmt.Printf("%s\n", result[start:])

}
