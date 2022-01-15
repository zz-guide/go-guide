package main

import (
	"fmt"
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
二叉搜索树中的插入操作

给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。

注意，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。 你可以返回 任意有效的结果 。
提示：

给定的树上的节点数介于 0 和 10^4 之间
每个节点都有一个唯一整数值，取值范围从 0 到 10^8
-10^8 <= val <= 10^8
新值和原始二叉搜索树中的任意节点值都不同


注意：1.插入的结果不止一种。 2.代码中的方式不会替换原来的结果，而是直接创建的方式


*/
func main() {
	searchNode := NewInsertSearchTreeNode()
	val := 5
	fmt.Println("二叉搜索树中的插入-递归法:", levelorder.TraversalRecursive(insertIntoBST(searchNode, val)))
	fmt.Println("二叉搜索树中的插入-迭代法:", levelorder.TraversalRecursive(insertIntoBST1(searchNode, val)))
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}

// insertIntoBST1 迭代法，利用搜索二叉树的性质进行遍历插入
func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	tempNode := &TreeNode{Val: val}
	if root == nil {
		return tempNode
	}

	cur := root
	for cur != nil {
		if val > cur.Val {
			if cur.Right == nil {
				cur.Right = tempNode
				break
			}

			cur = cur.Right
		}

		if val < cur.Val {
			if cur.Left == nil {
				cur.Left = tempNode
				break
			}

			cur = cur.Left
		}
	}

	return root
}
