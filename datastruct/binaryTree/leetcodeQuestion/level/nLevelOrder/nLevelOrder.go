package main

import (
	"fmt"
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/
N 叉树的层序遍历

给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。


*/
func main() {
	root := NewNode()

	fmt.Println("N叉树的层序遍历(迭代)：", nLevelOrder(root))
	fmt.Println("N叉树的层序遍历(递归)：", nLevelOrder2(root))
}

// levelOrderBottom N叉树层序遍历
func nLevelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var depth = 0 //深度,从0开始，最后结果是depth+1
	levelNodeQueue := []*Node{root}
	for len(levelNodeQueue) > 0 {
		res = append(res, []int{})
		length := len(levelNodeQueue)
		for j := 0; j < length; j++ {
			// 出栈，每次取第一个元素
			node := levelNodeQueue[0]
			levelNodeQueue = levelNodeQueue[1:]

			res[depth] = append(res[depth], node.Val)

			// children入队列
			for _, child := range node.Children {
				levelNodeQueue = append(levelNodeQueue, child)
			}
		}

		depth++
	}

	return res
}

func nLevelOrder2(root *Node) [][]int {
	var res [][]int
	var preorder func(node *Node, depth int)
	preorder = func(node *Node, depth int) {
		if node == nil {
			return
		}

		if depth == len(res) {
			res = append(res, []int{})
		}

		res[depth] = append(res[depth], node.Val)
		// children递归
		for _, child := range node.Children {
			preorder(child, depth+1)
		}

		return
	}

	preorder(root, 0)
	return res
}
