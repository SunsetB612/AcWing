package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	i, j := 0, 0
	for i < n && j < m {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j == m {
		return i - m
	}
	return -1
}

func main() {
	haystack := "leetcode"
	needle := "leeto"
	fmt.Println(strStr(haystack, needle))
}
