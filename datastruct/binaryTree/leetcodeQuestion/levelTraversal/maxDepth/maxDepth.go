package main

import (
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
二叉树的最大深度

给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明:叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度3 。
*/
func main() {
	root := NewNormalTree()

	log.Println("最大深度(递归法dfs)：", maxDepth(root))
	log.Println("最大深度(迭代法bfs)：", maxDepth1(root))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root} // 存储了每一层的全部结点
	var res = 0                // 默认就是1层
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 首次深度加一
			if i == 0 {
				res++
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		// 此处深度加一也可以
		// res++
	}

	return res
}
