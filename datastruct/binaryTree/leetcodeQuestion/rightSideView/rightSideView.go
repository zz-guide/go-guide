package main

import (
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/binary-tree-right-side-view/

二叉树的右视图

给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

提示:

二叉树的节点个数的范围是 [0,100]
-100 <= Node.val <= 100

注意：1.层序遍历即可，输出每一层最后一个元素的值

*/
func main() {
	root := NewNormalTree()
	log.Println("二叉树的右视图-层序遍历：", rightSideView(root))
}

// rightSideView 层序遍历
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int

	levelNodeQueue := []*TreeNode{root}
	for len(levelNodeQueue) > 0 {

		length := len(levelNodeQueue)
		for j := 0; j < length; j++ {
			// 出栈，每次取第一个元素
			node := levelNodeQueue[0]
			levelNodeQueue = levelNodeQueue[1:]

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
