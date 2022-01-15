package main

import (
	"fmt"
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/invert-binary-tree/
翻转一棵二叉树。

注意：就是左右互换
示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

*/
func main() {
	root := NewNormalTree()

	fmt.Println("翻转二叉树-递归:", invertTree(root))
	fmt.Println("翻转二叉树-迭代bfs", invertTree1(root))
	fmt.Println("翻转二叉树-迭代dfs", invertTree2(root))
}

// invertTree 递归法，前后序都可以,中序不行，可能会反转两次，相当于又回去了
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 前序
	//root.Left, root.Right = right, left
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	// 后序
	root.Left, root.Right = right, left
	return root
}

// invertTree1 迭代法，层序遍历
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	levelNodeQueue := []*TreeNode{root}
	for len(levelNodeQueue) > 0 {
		length := len(levelNodeQueue)
		for j := 0; j < length; j++ {
			// 出栈，每次取第一个元素
			node := levelNodeQueue[0]
			levelNodeQueue = levelNodeQueue[1:]

			node.Left, node.Right = node.Right, node.Left

			// 入队列，按照从左到右的顺序
			if node.Left != nil {
				levelNodeQueue = append(levelNodeQueue, node.Left)
			}

			if node.Right != nil {
				levelNodeQueue = append(levelNodeQueue, node.Right)
			}
		}
	}

	return root
}

// invertTree2 前序遍历迭代法dfs
func invertTree2(root *TreeNode) *TreeNode {
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
