package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	var b int
	fmt.Scan(&a, &b)

	if b == 0 {
		fmt.Println(0)
		return
	}

	n := len(a)
	carry := 0
	result := make([]byte, n)
	for i := n - 1; i >= 0; i-- {
		prod := int(a[i]-'0')*b + carry
		carry = prod / 10
		result[i] = byte(prod%10) + '0'
	}

	var prefix string
	if carry > 0 {
		prefix = strconv.Itoa(carry)
	}

	fmt.Printf("%s%s", prefix, result)
}
