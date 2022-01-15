package main

import (
	"fmt"
	. "go-guide/datastruct/binaryTree/treeNode"
	"sort"
)

/**
https://leetcode-cn.com/problems/count-complete-tree-nodes/
题目：给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~2h个节点

题解：1.二叉树的结点都可以通过遍历然后求出数量，前中后序遍历，层序遍历都可以
	2.题目给出的是完全二叉树，有一些特殊性质可以使用，完全二叉树不满的那颗子树一定是没有右节点。或者使用数组编号也可以减少搜索范围

*/
func main() {
	root := NewCompleteBinaryTree()

	fmt.Println("完全二叉树的结点个数-特性：", countNodes(root))
	fmt.Println("完全二叉树的结点个数-特性：", countNodes1(root))
}

// countNodes 利用完全二叉树特性的递归
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftH, rightH := 0, 0
	leftNode := root.Left

	rightNode := root.Right
	for leftNode != nil {
		leftNode = leftNode.Left
		leftH++
	}

	for rightNode != nil {
		rightNode = rightNode.Right
		rightH++
	}

	// 满二叉树
	if leftH == rightH {
		return (2 << leftH) - 1 // 2的k次方-1，k是树的深度
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// countNodes1 位运算+二分查找，时间复杂度：O(log^2 n)，空间复杂度：O(1) 利用数组编号来定位
func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}
	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			return false
		}
		bits := 1 << (level - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) - 1
}
