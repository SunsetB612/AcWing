package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*ListNode, len(vals))
	for i, v := range vals {
		nodes[i] = &ListNode{Val: v}
	}

	for i := 0; i < len(vals)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	return nodes[0]
}

func isPalindrome(head *ListNode) bool {
	var arr []int
	cur := head
	for cur.Next != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	arr = append(arr, cur.Val)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

func main() {
	head1 := buildList([]int{1, 2, 2, 1})
	fmt.Println(isPalindrome(head1))

	head2 := buildList([]int{1, 2})
	fmt.Println(isPalindrome(head2))
}
