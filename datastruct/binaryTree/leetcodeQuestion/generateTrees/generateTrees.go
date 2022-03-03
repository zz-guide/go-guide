package main

import (
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/unique-binary-search-trees-ii/

不同的二叉搜索树 II

给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。

提示：
1 <= n <= 8

*/
func main() {
	n := 2
	generateTrees(n)
}

// generateTrees 回溯
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	var helper func(start, end int) []*TreeNode
	helper = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}

		var allTrees []*TreeNode
		// 枚举可行根节点
		for i := start; i <= end; i++ {
			// 获得所有可行的左子树集合
			leftTrees := helper(start, i-1)
			// 获得所有可行的右子树集合
			rightTrees := helper(i+1, end)
			// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
			for _, left := range leftTrees {
				for _, right := range rightTrees {
					currTree := &TreeNode{Val: i}
					currTree.Left = left
					currTree.Right = right
					allTrees = append(allTrees, currTree)
				}
			}
		}

		return allTrees
	}

	return helper(1, n)
}
