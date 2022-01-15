package main

import (
	"fmt"
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/trim-a-binary-search-tree/
 修剪二叉搜索树

给你二叉搜索树的根节点 root ，同时给定最小边界low 和最大边界 high。通过修剪二叉搜索树，使得所有节点的值在[low, high]中。修剪树不应该改变保留在树中的元素的相对结构（即，如果没有被移除，原有的父代子代关系都应当保留）。 可以证明，存在唯一的答案。

所以结果应当返回修剪好的二叉搜索树的新的根节点。注意，根节点可能会根据给定的边界发生改变。

提示：

树中节点数在范围 [1, 104] 内
0 <= Node.val <= 104
树中每个节点的值都是唯一的
题目数据保证输入是一棵有效的二叉搜索树
0 <= low <= high <= 104

*/
func main() {
	root := NewNormalTree()
	low := 1
	high := 2
	fmt.Println("修剪二叉搜索树-递归：", trimBST(root, low, high))
	fmt.Println("修剪二叉搜索树-迭代：", trimBST1(root, low, high))
}

// trimBST dfs 前序遍历 递归
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if nil == root {
		return nil
	}

	for nil != root && (low > root.Val || high < root.Val) {
		// 整个左树都不要了
		if root.Val < low {
			root = root.Right
		} else if root.Val > high {
			root = root.Left
		}
	}

	if nil == root {
		return nil
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

// trimBST1 迭代法
func trimBST1(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	// 1.先处理跟结点
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else if root.Val > high {
			root = root.Left
		}
	}

	cur := root
	for cur != nil {
		// 处理左子树
		for cur.Left != nil && cur.Left.Val < low {
			cur.Left = cur.Left.Right
		}
		cur = cur.Left
	}

	cur = root
	for cur != nil {
		// 处理右子树
		for cur.Right != nil && cur.Right.Val > high {
			cur.Right = cur.Right.Left
		}
		cur = cur.Right
	}

	return root
}
