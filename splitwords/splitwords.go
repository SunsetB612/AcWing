package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	n := len(s)
	var words []string
	for i := 0; i < n; {
		j := i
		for j < n && s[j] != ' ' {
			j++
		}
		words = append(words, s[i:j])
		i = j + 1
	}

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWords("I am a student."))
}
