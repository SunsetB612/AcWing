package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	for len(a) < len(b) {
		a = a + "0"
	}
	for len(b) < len(a) {
		b = b + "0"
	}

	n := len(a)
	carry := 0
	result := make([]byte, n)

	for i := n - 1; i >= 0; i-- {
		sum := int(a[i]-'0') + int(b[i]-'0') + carry
		carry = sum / 10
		result[i] = byte(sum%10) + '0'
	}

	if carry > 0 {
		fmt.Printf("1%s", result)
	} else {
		fmt.Printf("%s", result)
	}
}
