package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	count := make(map[byte]int)
	ans := 0
	l := 0
	for r := 0; r < len(s); r++ {
		count[s[r]]++

		for count[s[r]] > 1 {
			count[s[l]]--
			l++
		}

		if r-l+1 > ans {
			ans = r - l + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
