package main

// 测试链接：https://leetcode.cn/problems/add-two-numbers/
func addTwoNumbers(head1 *ListNode, head2 *ListNode) *ListNode {
	var ans *ListNode = nil
	var cur *ListNode = nil
	carry := 0 //进位为0
	for head1 != nil || head2 != nil {
		sum := carry
		if head1 != nil {
			sum += head1.Val
		}

		if head2 != nil {
			sum += head2.Val
		}

		value := sum % 10
		carry = sum / 10
		if cur == nil {
			ans = &ListNode{Val: value, Next: nil}
			cur = ans
		} else {
			cur.Next = &ListNode{Val: value, Next: nil}
			cur = cur.Next
		}

		if head1 != nil {
			head1 = head1.Next
		}

		if head2 != nil {
			head2 = head2.Next
		}
	}

	if carry == 1 {
		cur.Next = &ListNode{Val: carry, Next: nil}
		cur = cur.Next
	}

	return ans
}
