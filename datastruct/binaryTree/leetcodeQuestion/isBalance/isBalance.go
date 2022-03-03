package main

import (
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
https://leetcode-cn.com/problems/balanced-binary-tree/
平衡二叉树
题目：输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

注意深度和高度的概念：
二叉树节点的深度：指从根节点到该节点的最长简单路径边的条数。
二叉树节点的高度：指从该节点到叶子节点的最长简单路径边的条数。


*/
func main() {
	root := NewBalanceBinaryTree()

	log.Println("是不是平衡二叉树-递归：", isBalanced(root))
}

// isBalanced 时间复杂度：O(n),空间复杂度：O(n)，自底向上
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftHeight := height(root.Left)
		rightHeight := height(root.Left)
		// 左树和右树有一个是-1，最终一定是不满足的
		if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
			return -1
		}

		return max(leftHeight, rightHeight) + 1
	}

	return height(root) > 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
