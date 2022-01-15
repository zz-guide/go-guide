package main

import (
	"fmt"
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/
在每个树行中找最大值

给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
*/
func main() {
	root := NewNormalTree()

	fmt.Println("每一层最右侧结点-迭代：", largestValues(root))
	fmt.Println("每一层最右侧结点-递归", largestValues2(root))
}

// largestValues 层序遍历,遍历每一层求最大值
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	levelNodeQueue := []*TreeNode{root}

	// math.MinInt,leetcode有时候没有引入math包
	intSize := 32 << (^uint(0) >> 63) // 32 or 64
	MinInt := -1 << (intSize - 1)

	for len(levelNodeQueue) > 0 {
		length := len(levelNodeQueue)

		var max = MinInt
		for j := 0; j < length; j++ {
			node := levelNodeQueue[0]
			levelNodeQueue = levelNodeQueue[1:]

			// 求最大值
			if node.Val >= max {
				max = node.Val
			}

			// 入队列，按照从左到右的顺序
			if node.Left != nil {
				levelNodeQueue = append(levelNodeQueue, node.Left)
			}

			if node.Right != nil {
				levelNodeQueue = append(levelNodeQueue, node.Right)
			}
		}

		res = append(res, max)
	}

	return res
}

// largestValues2 递归的过程没法判断当前是不是最后一个结点，只能层序遍历完再把最后一个元素取出来
func largestValues2(root *TreeNode) []int {
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

	var ret []int
	intSize := 32 << (^uint(0) >> 63) // 32 or 64
	MinInt := -1 << (intSize - 1)

	for _, item := range res {

		max := MinInt
		for _, item2 := range item {
			if item2 >= max {
				max = item2
			}
		}

		ret = append(ret, max)
	}

	return ret
}
