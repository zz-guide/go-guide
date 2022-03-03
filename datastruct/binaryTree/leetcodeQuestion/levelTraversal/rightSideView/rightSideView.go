package main

import (
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/binary-tree-right-side-view/
二叉树的右视图

给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

示例 1:

输入:[1,2,3,null,5,null,4]
输出:[1,3,4]
示例 2:

输入:[1,null,3]
输出:[1,3]
示例 3:

输入:[]
输出:[]

注意：其实就是返回每一层的最又侧结点
*/
func main() {
	root := NewNormalTree()

	log.Println("每一层最右侧结点-迭代：", rightSideView(root))
	log.Println("每一层最右侧结点-递归", rightSideView2(root))
}

// rightSideView 层序遍历，每次只添加每一层的最后一个元素
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	levelNodeQueue := []*TreeNode{root}
	for len(levelNodeQueue) > 0 {
		length := len(levelNodeQueue)
		for j := 0; j < length; j++ {
			node := levelNodeQueue[0]
			levelNodeQueue = levelNodeQueue[1:]

			// 每次只添加最后一个元素的值到结果集
			if j == length-1 {
				res = append(res, node.Val)
			}

			// 入队列，按照从左到右的顺序
			if node.Left != nil {
				levelNodeQueue = append(levelNodeQueue, node.Left)
			}

			if node.Right != nil {
				levelNodeQueue = append(levelNodeQueue, node.Right)
			}
		}
	}

	return res
}

// rightSideView2 递归的过程没法判断当前是不是最后一个结点，只能层序遍历完再把最后一个元素取出来
func rightSideView2(root *TreeNode) []int {
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
	for _, item := range res {
		ret = append(ret, item[len(item)-1])
	}

	return ret
}
