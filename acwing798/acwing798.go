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

	var n, m, q int
	fmt.Fscan(reader, &n, &m, &q)

	arr := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fscan(reader, &arr[i][j])
		}
	}

	diff := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		diff[i] = make([]int, m+2)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			diff[i][j] = arr[i][j] - arr[i-1][j] - arr[i][j-1] + arr[i-1][j-1]
		}
	}

	for i := 0; i < q; i++ {
		var x1, y1, x2, y2, c int
		fmt.Fscan(reader, &x1, &y1, &x2, &y2, &c)
		diff[x1][y1] += c
		diff[x1][y2+1] -= c
		diff[x2+1][y1] -= c
		diff[x2+1][y2+1] += c
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			diff[i][j] = diff[i-1][j] + diff[i][j-1] - diff[i-1][j-1] + diff[i][j]
			fmt.Fprintf(writer, "%d ", diff[i][j])
		}
		fmt.Fprintln(writer)
	}
}
