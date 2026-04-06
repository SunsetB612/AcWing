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

	prefix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		prefix[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			var num int
			fmt.Fscan(reader, &num)
			prefix[i][j] = prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1] + num
		}
	}

	for i := 0; i < q; i++ {
		var x1, y1, x2, y2 int
		fmt.Fscan(reader, &x1, &y1, &x2, &y2)
		ans := prefix[x2][y2] - prefix[x1-1][y2] - prefix[x2][y1-1] + prefix[x1-1][y1-1]
		fmt.Fprintln(writer, ans)
	}
}
