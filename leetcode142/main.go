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

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

func main() {
	head1 := buildList([]int{3, 2, 0, -4}, 1)
	node := detectCycle(head1)
	if node != nil {
		fmt.Println(node.Val)
	}
}
