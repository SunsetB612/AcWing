package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n, m, x int
	fmt.Fscan(reader, &n, &m, &x)

	a := make([]int, n)
	b := make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	for j := 0; j < m; j++ {
		fmt.Fscan(reader, &b[j])
	}

	i, j := 0, m-1
	for i < n && j >= 0 {
		sum := a[i] + b[j]
		if sum > x {
			j--
		} else if sum < x {
			i++
		} else {
			fmt.Fprintln(writer, i, j)
			break
		}
	}
}
