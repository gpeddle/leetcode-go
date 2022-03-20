package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummyHead := ListNode{}
	p := l1
	q := l2
	curr := &dummyHead
	carry := 0

	for p.Next != nil || q.Next != nil {
		var x int
		if p.Next != nil {
			x = p.Val
		} else {
			x = 0
		}

		var y int
		if q.Next != nil {
			y = q.Val
		} else {
			y = 0
		}

		sum := carry + x + y

		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next

		if p.Next != nil {
			p = p.Next
		}
		if q.Next != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next
}
