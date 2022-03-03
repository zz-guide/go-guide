package main

import (
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/sum-of-left-leaves/
左叶子结点之和

计算给定二叉树的所有左叶子之和。

注意：左叶子概念:如果左节点不为空，且左节点没有左右孩子，那么这个节点就是左叶子

*/
func main() {
	root := NewNormalTree()

	log.Println("左叶结点之和-递归：", sumOfLeftLeaves(root))
	log.Println("左叶结点之和-递归：", sumOfLeftLeaves1(root))
}

// sumOfLeftLeaves dfs 前序遍历 递归
func sumOfLeftLeaves(root *TreeNode) int {
	var res int

	// 前序遍历
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		// 如果是左叶结点，累加
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			res += node.Left.Val
		}

		if node.Left != nil {
			travel(node.Left)
		}

		if node.Right != nil {
			travel(node.Right)
		}
	}

	travel(root)
	return res
}

// sumOfLeftLeaves1 前序遍历 dfs迭代
func sumOfLeftLeaves1(root *TreeNode) int {
	var res int

	var stack []*TreeNode

	if root != nil {
		stack = append(stack, root)
	}

	for len(stack) > 0 {
		l := len(stack)

		node := stack[l-1]
		stack = stack[:l-1]

		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			res += node.Left.Val
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return res
}
