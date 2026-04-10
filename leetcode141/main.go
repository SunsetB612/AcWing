package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(vals []int, pos int) *ListNode {
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

	if pos >= 0 {
		nodes[len(vals)-1].Next = nodes[pos]
	}
	return nodes[0]
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}

	return false
}

func main() {
	head1 := buildList([]int{3, 2, 0, -4}, 1)
	fmt.Println(hasCycle(head1))

	head2 := buildList([]int{1, 2}, 0)
	fmt.Println(hasCycle(head2))

	head3 := buildList([]int{1}, -1)
	fmt.Println(hasCycle(head3))

}
