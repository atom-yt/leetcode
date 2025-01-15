package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preTraverse(t *TreeNode) {
	if t == nil {
		return
	}
	fmt.Println(t.Val)
	preTraverse(t.Left)
	preTraverse(t.Right)
}

func midTraverse(t *TreeNode) {
	if t == nil {
		return
	}
	midTraverse(t.Left)
	fmt.Println(t.Val)
	midTraverse(t.Right)
}

func backTraverse(t *TreeNode) {
	if t == nil {
		return
	}
	backTraverse(t.Left)
	backTraverse(t.Right)
	fmt.Println(t.Val)
}

func levelTraverse(t *TreeNode) {
	if t == nil {
		return
	}

	queue := []*TreeNode{t}

	for len(queue) > 0 {
		current := queue[0]

		queue = queue[1:]
		fmt.Println(current.Val)
		
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}

func main() {
	t := &TreeNode{Val: 1}
	t.Left = &TreeNode{Val: 2}
	t.Right = &TreeNode{Val: 3}
	t.Left.Left = &TreeNode{Val: 4}
	t.Left.Right = &TreeNode{Val: 5}
	t.Right.Right = &TreeNode{Val: 6}
	
	levelTraverse(t)

}