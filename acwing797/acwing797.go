package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func addNum(arr []int, l, r, c int) {
	for i := l; i <= r; i++ {
		arr[i] += c
	}
}

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	diff := make([]int, n+2)
	for i := 1; i <= n; i++ {
		diff[i] = arr[i] - arr[i-1]
	}

	var l, r, c int
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &l, &r, &c)
		diff[l] += c
		diff[r+1] -= c
	}

	for i := 1; i <= n; i++ {
		diff[i] += diff[i-1]
		fmt.Fprintf(writer, "%d ", diff[i])
	}
}
