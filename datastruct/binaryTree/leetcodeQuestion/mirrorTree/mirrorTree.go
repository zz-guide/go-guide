package main

import (
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/

二叉树的镜像

请完成一个函数，输入一个二叉树，该函数输出它的镜像。

限制：

0 <= 节点个数 <= 1000

注意：本题与主站 226 题相同：https://leetcode-cn.com/problems/invert-binary-tree/


*/
func main() {
	root1 := NewNormalTree()
	root2 := NewNormalTree()
	root3 := NewNormalTree()

	log.Println("二叉树的镜像-递归：", levelorder.TraversalRecursive(mirrorTree(root1)))
	log.Println("二叉树的镜像-层序遍历：", levelorder.TraversalRecursive(mirrorTree1(root2)))
	log.Println("二叉树的镜像-前序遍历：", levelorder.TraversalRecursive(mirrorTree2(root3)))
}

// mirrorTree 递归
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := mirrorTree(root.Left)
	right := mirrorTree(root.Right)

	root.Left = right
	root.Right = left

	return root
}

// mirrorTree1 层序遍历
func mirrorTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}

// mirrorTree2 前序遍历迭代法dfs
func mirrorTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var stack []*TreeNode
	node := root

	for node != nil || len(stack) > 0 {
		for node != nil {
			// 此处交换左右树
			node.Left, node.Right = node.Right, node.Left
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}

	return root
}
