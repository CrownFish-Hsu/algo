package main

func mergeTwoLinkList(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}

	head := &ListNode{}
	cur1 := &ListNode{}
	cur2 := &ListNode{}
	if head1.Val <= head2.Val {
		head = head1
		cur1 = head.Next
		cur2 = head2
	} else {
		head = head2
		cur1 = head.Next
		cur2 = head1
	}

	prev := head
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			prev.Next = cur1
			cur1 = cur1.Next
		} else {
			prev.Next = cur2
			cur2 = cur2.Next
		}

		prev = prev.Next
	}

	if cur1 != nil {
		prev.Next = cur1
	} else {
		prev.Next = cur2
	}

	return head
}
