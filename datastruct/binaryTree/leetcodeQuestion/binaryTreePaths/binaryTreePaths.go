package main

import (
	"fmt"
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
	"strconv"
)

/**
题目：https://leetcode-cn.com/problems/binary-tree-paths/
二叉树的所有路径
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。

*/
func main() {
	root := NewNormalTree()

	log.Println("二叉树所有路径-递归：", binaryTreePaths(root))
	log.Println("二叉树所有路径-递归：", binaryTreePaths1(root))
}

// binaryTreePaths dfs 前序遍历 递归
func binaryTreePaths(root *TreeNode) []string {
	var res []string

	// 前序遍历
	var travel func(node *TreeNode, s string)
	travel = func(node *TreeNode, s string) {
		// 如果当前结点是叶子结点，加入结果集
		if node.Left == nil && node.Right == nil {
			v := s + strconv.Itoa(node.Val)
			res = append(res, v)
			return
		}

		// 先把当前结点的值串到s上
		s = s + strconv.Itoa(node.Val) + "->"

		if node.Left != nil {
			travel(node.Left, s)
		}

		if node.Right != nil {
			travel(node.Right, s)
		}
	}

	travel(root, "")
	return res
}

// binaryTreePaths1 前序遍历 dfs迭代
func binaryTreePaths1(root *TreeNode) []string {
	var res []string

	var stack []*TreeNode
	var paths []string // 与上边栈一一对应，存当前结点的路径

	if root != nil {
		stack = append(stack, root)
		paths = append(paths, "")
	}

	for len(stack) > 0 {
		l := len(stack)

		node := stack[l-1]
		stack = stack[:l-1]

		path := paths[l-1]
		paths = paths[:l-1]

		if node.Left == nil && node.Right == nil {
			res = append(res, path+strconv.Itoa(node.Val))
			continue
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
	}

	return res
}
