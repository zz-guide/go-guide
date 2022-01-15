package main

import (
	"fmt"
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
层序遍历自底向上

给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层序遍历为：

[
  [15,7],
  [9,20],
  [3]
]
*/
func main() {
	root := NewNormalTree()

	fmt.Println("层序遍历迭代法(自底向上)：", levelOrderBottom(root))
	fmt.Println("层序遍历递归法(自底向上)：", levelOrderBottom2(root))
}

// levelOrderBottom 层序遍历，然后取反
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var depth = 0 //深度,从0开始，最后结果是depth+1
	levelNodeQueue := []*TreeNode{root}
	for len(levelNodeQueue) > 0 {
		// 进入新的一层需要留出一个位置放结果
		res = append(res, []int{})
		length := len(levelNodeQueue)
		for j := 0; j < length; j++ {
			// 出栈，每次取第一个元素
			node := levelNodeQueue[0]
			levelNodeQueue = levelNodeQueue[1:]

			res[depth] = append(res[depth], node.Val)

			// 入队列，按照从左到右的顺序
			if node.Left != nil {
				levelNodeQueue = append(levelNodeQueue, node.Left)
			}

			if node.Right != nil {
				levelNodeQueue = append(levelNodeQueue, node.Right)
			}
		}

		depth++
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}

func levelOrderBottom2(root *TreeNode) [][]int {
	var res [][]int
	var preorder func(node *TreeNode, depth int)
	preorder = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		// 很巧妙，增加位置的时机只在depth == len(res)，并且是在判空之前
		if depth == len(res) {
			res = append(res, []int{})
		}

		res[depth] = append(res[depth], node.Val)
		preorder(node.Left, depth+1)
		preorder(node.Right, depth+1)
		return
	}

	preorder(root, 0)

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
