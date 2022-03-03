package main

import (
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/***
二叉搜索树的最小绝对差：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

差值是一个正数，其数值等于两值之差的绝对值。
提示：

树中节点的数目范围是 [2, 104]
0 <= Node.val <= 105

注意：1.值都是大于0的 2.中序遍历之后的数组，最小绝对差一定是连续的2个结点，可能在中间，也可能在两边

*/
func main() {
	root := NewNormalTree()
	log.Println("二叉搜索树的最小绝对差-递归：", getMinimumDifference(root))
}

// getMinimumDifference 递归法，需要注意的是，得保存一个pre结点的值，只靠递归的话无法知道前一个结点是谁
func getMinimumDifference(root *TreeNode) int {
	intSize := 32 << (^uint(0) >> 63) // 32 or 64
	var res int = 1<<(intSize-1) - 1

	var preNode *TreeNode
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		helper(node.Left)
		if preNode != nil {
			res = min(res, node.Val-preNode.Val)
		}

		preNode = node
		helper(node.Right)
	}

	helper(root)
	return res
}

func min(x int, y int) int {
	if x < y {
		return x
	}

	return y
}

// getMinimumDifference1 迭代法
func getMinimumDifference1(root *TreeNode) int {
	intSize := 32 << (^uint(0) >> 63) // 32 or 64
	var res int = 1<<(intSize-1) - 1

	var preNode *TreeNode

	var stack []*TreeNode
	node := root
	for node != nil || 0 < len(stack) {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 中序

		if preNode != nil {
			res = min(res, node.Val-preNode.Val)
		}

		preNode = node
		node = node.Right
	}

	return res
}
