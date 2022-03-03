package main

import (
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/pOCWxh/

二叉树剪枝

给定一个二叉树 根节点root，树的每个节点的值要么是 0，要么是 1。请剪除该二叉树中所有节点的值为 0 的子树。

节点 node 的子树为node 本身，以及所有 node的后代。

提示:

二叉树的节点个数的范围是 [1,200]
二叉树节点的值只会是 0 或 1

注意:
1.二叉树的结点值只能是0或者1
2.值是0的叶子结点需要删除
3.左子树和右子树有一个1就不能被删除
4.移除的子树中包含所有之后的结点
*/
func main() {
	root := NewNormalTree3()
	log.Println("二叉树剪枝-递归：", levelorder.TraversalRecursive(pruneTree(root)))
	log.Println("二叉树剪枝-迭代：", levelorder.TraversalRecursive(pruneTree1(root)))
}

// pruneTree 递归实现
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var prune func(node *TreeNode) *TreeNode
	prune = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}

		node.Left = prune(node.Left)
		node.Right = prune(node.Right)

		if node.Left == nil && node.Right == nil && node.Val == 0 {
			return nil
		}

		return node
	}

	root = prune(root)
	return root
}

// pruneTree1 后序遍历迭代
func pruneTree1(root *TreeNode) *TreeNode {
	var stack []*TreeNode
	var pre *TreeNode

	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		if node.Right != nil && node.Right != pre {
			node = node.Right
		} else {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if node.Left == nil && node.Right == nil && node.Val == 0 {
				if len(stack) > 0 {
					tmpNode := stack[len(stack)-1]
					if tmpNode.Left == node {
						tmpNode.Left = nil
					} else if tmpNode.Right == node {
						tmpNode.Right = nil
					}
				} else {
					return nil
				}
			}

			pre = node
			node = nil
		}
	}

	return root
}
