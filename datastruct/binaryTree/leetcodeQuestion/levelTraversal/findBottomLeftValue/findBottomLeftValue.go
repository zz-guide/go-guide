package main

import (
	"container/list"
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/find-bottom-left-tree-value/
找树左下角的值
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
假设二叉树中至少有一个节点。

注意：重点留意最左边结点的含义，理解为最底层从左数第一个结点。最底层就是深度最大，求左边第一个那肯定就是前序遍历了
*/
func main() {
	root := NewNormalTree()

	log.Println("最左边的值(前序遍历递归法)：", findBottomLeftValue(root))
	log.Println("最左边的值(层序遍历迭代法)：", findBottomLeftValue1(root))
}

// findBottomLeftValue 递归解法，前序遍历，dfs
func findBottomLeftValue(root *TreeNode) int {
	var maxDeep int // 深度
	var value int   // 结果

	// 当只有一个结点的时候，返回自身的值，题目中特别强调
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	// 根左右，前序遍历
	var findLeftValue func(root *TreeNode, deep int)
	findLeftValue = func(root *TreeNode, deep int) {
		if root == nil {
			return
		}

		// 递归到叶子节点
		if root.Left == nil && root.Right == nil {
			if deep > maxDeep {
				// 此处直接保存结果，不关心顺序
				value = root.Val
				maxDeep = deep
			}

			return
		}

		findLeftValue(root.Left, deep+1)
		findLeftValue(root.Right, deep+1)
	}

	// 从root开始递归，初始深度为0
	findLeftValue(root, maxDeep)
	return value
}

// findBottomLeftValue1 迭代法，层序遍历，bfs
func findBottomLeftValue1(root *TreeNode) int {
	queue := list.New()
	var res int
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			// 先进先出队列
			node := queue.Remove(queue.Front()).(*TreeNode)
			// 记录本层第一个值，
			if i == 0 {
				res = node.Val
			}

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

	return res
}
