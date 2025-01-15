package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Children []*TreeNode
}

func preOrderTravers(root *TreeNode) {
	if root == nil {
		return 
	}
	fmt.Println(root.Val)
	for _, i := range root.Children {
		preOrderTravers(i)
	}
}

func postOrderTravers(root *TreeNode) {
	if root == nil {
		return 
	}
	
	for _, i := range root.Children {
		postOrderTravers(i)
	}
	fmt.Println(root.Val)
}

func levelOrderTravers(root *TreeNode) {
	if root == nil {
		return 
	}
	
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		current := queue[0]
		fmt.Println(current.Val)
		queue = queue[1:]

		for _, i := range current.Children {
			queue = append(queue, i)
		}
	}
}

func main() {
	t := &TreeNode{Val: 1}
	t1 := &TreeNode{Val: 2}
	t2 := &TreeNode{Val: 3}
	t3 := &TreeNode{Val: 4}
	t.Children = []*TreeNode{t1, t2, t3}
	postOrderTravers(t)
	
}