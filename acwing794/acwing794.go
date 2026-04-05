package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	var b int
	fmt.Scan(&a, &b)

	n := len(a)
	remainder := 0
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		remainder = remainder*10 + int(a[i]-'0')
		result[i] = byte(remainder/b) + '0'
		remainder = remainder % b
	}

	start := 0
	for start < n-1 && result[start] == '0' {
		start++
	}

	fmt.Printf("%s\n", result[start:])
	fmt.Println(strconv.Itoa(remainder))
}
