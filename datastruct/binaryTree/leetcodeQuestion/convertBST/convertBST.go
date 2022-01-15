package main

import (
	"fmt"
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
把二叉搜索树转换为累加树

给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），使每个节点 node的新值等于原树中大于或等于node.val的值之和。

提醒一下，二叉搜索树满足下列约束条件：

节点的左子树仅包含键 小于 节点键的节点。
节点的右子树仅包含键 大于 节点键的节点。
左右子树也必须是二叉搜索树。
注意：本题和 1038:https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree/ 相同


注意：1.对于搜索二叉树，右边的值最大，所以从右往左查找比较好 2.遍历顺序，由根左

*/
func main() {
	root := NewSearchTreeNode()
	fmt.Println("二叉搜索树转换为累加树-递归：", levelorder.TraversalRecursive(convertBST(root)))
	root1 := NewSearchTreeNode()
	fmt.Println("二叉搜索树转换为累加树-迭代：", levelorder.TraversalRecursive(convertBST1(root1)))
}

// 递归法
func convertBST(root *TreeNode) *TreeNode {
	var sum int
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}

		helper(root.Right)
		sum += root.Val
		root.Val = sum
		helper(root.Left)
	}

	helper(root)
	return root
}

// 迭代法
func convertBST1(root *TreeNode) *TreeNode {
	var stack []*TreeNode
	node := root
	var sum int

	for node != nil || 0 < len(stack) {
		for node != nil {
			stack = append(stack, node)
			node = node.Right
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 中序的反顺序
		sum += node.Val
		node.Val = sum
		node = node.Left
	}

	return root
}
