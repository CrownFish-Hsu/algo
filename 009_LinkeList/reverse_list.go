package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode = nil
	var next *ListNode = nil

	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}
