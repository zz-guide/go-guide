package main

import (
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
给定一棵树的前序遍历 preorder 与中序遍历  inorder。请构造二叉树并返回其根节点。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
*/
func main() {
	preorder := []int{9, 3, 15, 20, 7}
	inorder := []int{9, 15, 7, 20, 3}
	log.Println("中序后序构造二叉树-递归：", levelorder.TraversalRecursive(buildTree(preorder, inorder)))
	log.Println("中序后序构造二叉树-迭代：", levelorder.TraversalRecursive(buildTree1(preorder, inorder)))
}

// buildTree 递归法构建，以根节点为切分，将左右区间递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

// buildTree1 迭代法
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	var stack []*TreeNode
	stack = append(stack, root)

	var inorderIndex int
	for i := 1; i < len(preorder); i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{Val: preorderVal}
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{Val: preorderVal}
			stack = append(stack, node.Right)
		}
	}

	return root
}
