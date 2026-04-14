package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})

		key := string(b)
		m[key] = append(m[key], str)
	}

	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))

	strs = []string{""}
	fmt.Println(groupAnagrams(strs))

	strs = []string{"a"}
	fmt.Println(groupAnagrams(strs))
}
