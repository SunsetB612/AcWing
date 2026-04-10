package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var res []rune
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			res = append(res, c)
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		if res[i] != res[j] {
			return false
		}
	}
	return true
}

func main() {
	s := "0P"
	fmt.Println(isPalindrome(s))
}
