package main

import (
	"fmt"
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
给定一个二叉树，计算 整个树 的坡度 。

一个树的 节点的坡度 定义即为，该节点左子树的节点之和和右子树节点之和的 差的绝对值 。如果没有左子树的话，左子树的节点之和为 0 ；没有右子树的话也是一样。空结点的坡度是 0 。

整个树 的坡度就是其所有节点的坡度之和。

*/
func main() {
	root := NewNormalTree()

	fmt.Println("是不是对称二叉树-递归:", findTilt(root))
}

// findTilt 递归，后序遍历
func findTilt(root *TreeNode) int {
	var res int
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		sumLeft := dfs(node.Left)
		sumRight := dfs(node.Right)
		// 计算整颗树的坡度，坡度和可能是负的
		res += abs(sumLeft - sumRight)
		// 返回当前结点的坡度
		return sumLeft + sumRight + node.Val
	}

	dfs(root)
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
