package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	count := make(map[int]int)
	ans := 0
	l := 0
	for r := 0; r < n; r++ {
		count[arr[r]]++
		for count[arr[r]] > 1 {
			count[arr[l]]--
			l++
		}

		if r-l+1 > ans {
			ans = r - l + 1
		}
	}
	fmt.Println(ans)

}
