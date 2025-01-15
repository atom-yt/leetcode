package main

/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
 

示例 1：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

示例 2：
输入：head = [1,2]
输出：[2,1]

示例 3：
输入：head = []
输出：[]

dummy -> 1 -> 2 -> 3
 pre
 current

提示：

链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000
 
*/

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseLinkNode(root *ListNode) *ListNode {
	var pre *ListNode
	current := root


	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}

	return pre

}

func printLinkNode(root *ListNode) {
	current := root
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func main() {
	t := &ListNode{Val: 1}
	t.Next = &ListNode{Val: 2}
	t.Next.Next = &ListNode{Val: 3}
	t.Next.Next.Next = &ListNode{Val: 4}

	printLinkNode(reverseLinkNode(t))

}