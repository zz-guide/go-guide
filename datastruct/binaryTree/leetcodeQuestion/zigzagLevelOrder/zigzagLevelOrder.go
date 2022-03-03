package main

import (
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

二叉树的锯齿形层序遍历

给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[3],[20,9],[15,7]]
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]

提示：

树中节点数目在范围 [0, 2000] 内
-100 <= Node.val <= 100

*/
func main() {
	root := NewNormalTree()
	log.Println("二叉树的锯齿形层序遍历-BFS迭代：", zigzagLevelOrder(root))
}

// zigzagLevelOrder 迭代bfs
// 只能是先遍历完当前层然后反转，因为在遍历的过程中stack中元素会变化，不太好处理
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	levelNodeQueue := []*TreeNode{root}

	level := 1 // 表示上一层是第几层
	for len(levelNodeQueue) != 0 {
		var temp []int
		length := len(levelNodeQueue)
		for i := 0; i < length; i++ {
			node := levelNodeQueue[0]
			levelNodeQueue = levelNodeQueue[1:]
			temp = append(temp, node.Val)

			if node.Left != nil {
				levelNodeQueue = append(levelNodeQueue, node.Left)
			}

			if node.Right != nil {
				levelNodeQueue = append(levelNodeQueue, node.Right)
			}
		}

		// 偶数层反转
		if level%2 == 0 {
			for j, k := 0, len(temp)-1; j < k; j, k = j+1, k-1 {
				temp[j], temp[k] = temp[k], temp[j]
			}
		}

		res = append(res, temp)
		level++
	}

	return res
}
