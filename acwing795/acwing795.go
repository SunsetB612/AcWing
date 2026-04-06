package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&prefix[i])
		prefix[i] += prefix[i-1]
	}

	var l, r int
	for i := 0; i < m; i++ {
		fmt.Scan(&l, &r)
		fmt.Println(prefix[r] - prefix[l-1])
	}
}
