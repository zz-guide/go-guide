package main

import (
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/
二叉树的层平均值

给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。


示例 1：

输入：
    3
   / \
  9  20
    /  \
   15   7
输出：[3, 14.5, 11]
解释：
第 0 层的平均值是 3 ,  第1层是 14.5 , 第2层是 11 。因此返回 [3, 14.5, 11] 。

提示：

节点值的范围在32位有符号整数范围内。
*/
func main() {
	root := NewNormalTree()

	log.Println("每一层的平均值-迭代：", averageOfLevels(root))
	log.Println("每一层的平均值-递归", averageOfLevels2(root))
}

// averageOfLevels 层序遍历，计算平均值
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	var res []float64
	levelNodeQueue := []*TreeNode{root}
	for len(levelNodeQueue) > 0 {
		length := len(levelNodeQueue)

		var sum int
		for j := 0; j < length; j++ {
			node := levelNodeQueue[0]
			levelNodeQueue = levelNodeQueue[1:]

			sum += node.Val

			if node.Left != nil {
				levelNodeQueue = append(levelNodeQueue, node.Left)
			}

			if node.Right != nil {
				levelNodeQueue = append(levelNodeQueue, node.Right)
			}
		}

		res = append(res, float64(sum)/float64(length))
	}

	return res
}

// averageOfLevels2 递归法，最后计算平均值
func averageOfLevels2(root *TreeNode) []float64 {
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

	var ret []float64
	for _, item := range res {

		var sum int
		for _, item2 := range item {
			sum += item2
		}

		ret = append(ret, float64(sum)/float64(len(item)))
	}

	return ret
}
