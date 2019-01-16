package main

import (
	`fmt`
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Shownode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var re *ListNode
	re = new(ListNode)
	node := re
	re.Next = node
	re.Val = 0

	var carry int
	for l1 != nil || l2 != nil {
		var x, y, sum int
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		sum = x + y + carry
		carry = sum / 10
		node.Next = &ListNode{Val: sum % 10}
		node = node.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry == 1 {
		node.Next = &ListNode{Val: carry}
		node = node.Next
	}
	return re.Next
}
