package leetcode

import (
	"math"
)

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

	n1 := toInteger(l1)
	if n1 == 0 {
		return l2
	}
	n2 := toInteger(l2)
	if n2 == 0 {
		return l1
	}

	sum := n1 + n2

	result := toList(sum)
	return result
}

func toInteger(list *ListNode) int {

	var result = 0
	pow := 0
	for list != nil {
		result += list.Val * int(math.Pow10(pow))
		pow++
		list = list.Next
	}
	return result
}

func toList(num int) *ListNode {

	var (
		head     *ListNode
		current  *ListNode
		previous *ListNode
	)
	working := num

	for working != 0 {

		digit := working % 10
		working = (working - digit) / 10
		current = &ListNode{
			Val:  digit,
			Next: nil,
		}

		if head == nil {
			head = current
		} else {
			previous.Next = current
		}
		previous = current

	}
	return head
}

/*
func reverse(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	working := head.Next
	reversed := head
	reversed.Next = nil

	for working != nil {
		tmp := working
		working = working.Next
		tmp.Next = reversed
		reversed = tmp
	}
	return reversed
}
*/
