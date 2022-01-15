package levelorder

import (
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
https://leetcode-cn.com/problems/binary-tree-level-order-traversal/solution/bfs-de-shi-yong-chang-jing-zong-jie-ceng-xu-bian-l/
层序遍历
示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

注意：只返回有效的结果，nil不需要存
*/

// TraversalRecursive 利用二叉树的先序遍历（根左右），记录二叉树深度。层次递归遍历，每次进入下层则深度+1，若当前深度大于等于当前res的长度，则给res扩容。因为是先序根左右，符合层序遍历的层级次序。
func TraversalRecursive(root *TreeNode) [][]int {
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
	return res
}

// TraversalQueue BFS,队列思想
func TraversalQueue(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	var depth = 0 //深度,从0开始，最后结果是depth+1
	levelNodeQueue := []*TreeNode{root}
	for len(levelNodeQueue) > 0 {
		// 进入新的一层需要留出一个位置放结果
		ret = append(ret, []int{})
		length := len(levelNodeQueue)
		for j := 0; j < length; j++ {
			// 出栈，每次取第一个元素
			node := levelNodeQueue[0]
			levelNodeQueue = levelNodeQueue[1:]

			ret[depth] = append(ret[depth], node.Val)

			// 入队列，按照从左到右的顺序
			if node.Left != nil {
				levelNodeQueue = append(levelNodeQueue, node.Left)
			}

			if node.Right != nil {
				levelNodeQueue = append(levelNodeQueue, node.Right)
			}
		}

		depth++
	}

	return ret
}
