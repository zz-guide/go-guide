package main

import (
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/***
路径总和：https://leetcode-cn.com/problems/path-sum/
给你二叉树的根节点root 和一个表示目标和的整数targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和targetSum 。如果存在，返回 true ；否则，返回 false 。

叶子节点 是指没有子节点的节点
*/
func main() {
	root := NewNormalTree()

	targetSum := 11
	log.Println("路径总和-递归：", hasPathSum(root, targetSum))
	log.Println("路径总和-递归：", hasPathSum1(root, targetSum))
}

// hasPathSum 递归法
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum-root.Val == 0
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// hasPathSum1 迭代法BFS,前序遍历
func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	queue := []*TreeNode{root}
	pathSum := []int{root.Val}
	for len(queue) != 0 {
		n := queue[0]
		p := pathSum[0]

		queue = queue[1:]
		pathSum = pathSum[1:]

		// 判断:当遍历到叶子节点时,如果p等于给定的值就返回真
		if n.Left == nil && n.Right == nil {
			if p == targetSum {
				return true
			}

			continue
		}

		if n.Left != nil {
			queue = append(queue, n.Left)
			pathSum = append(pathSum, n.Left.Val+p)
		}

		if n.Right != nil {
			queue = append(queue, n.Right)
			pathSum = append(pathSum, n.Right.Val+p)
		}
	}

	return false
}
