package main

import (
	"fmt"
	. "go-guide/datastruct/binaryTree/treeNode"
	"math"
)

/**
https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/
二叉树的最小深度

给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。


注意：当left和right都为nil，说明是最低点，其次从root开始遍历，就是最低的
*/
func main() {
	root := NewNormalTree()

	fmt.Println("最小深度(递归法dfs)：", minDepth(root))
	fmt.Println("最小深度(迭代法bfs)：", minDepth1(root))
}
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}
	return minD + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root} // 存储了每一层的全部结点
	var res = 0                // 默认就是1层

Loop:
	for len(queue) != 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 首次深度加一
			if i == 0 {
				res++
			}

			if cur.Left == nil && cur.Right == nil {
				break Loop
			}

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}

	return res
}
