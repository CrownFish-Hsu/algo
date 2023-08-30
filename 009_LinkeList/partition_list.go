package main

// https://leetcode.cn/problems/partition-list/
func partition(head *ListNode, x int) *ListNode {
	var next, ltHead, ltTail, gtHead, gtTail *ListNode = nil, nil, nil, nil, nil

	for head != nil {
		next = head.Next
		head.Next = nil
		if head.Val < x {
			if ltHead == nil {
				ltHead = head
			} else {
				ltTail.Next = head
			}
			ltTail = head
		} else {
			if gtHead == nil {
				gtHead = head
			} else {
				gtTail.Next = head
			}
			gtTail = head
		}

		head = next
	}

	if ltHead == nil {
		return gtHead
	}

	ltTail.Next = gtHead
	return ltHead
}
