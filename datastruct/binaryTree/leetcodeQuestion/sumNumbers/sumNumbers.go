package main

import (
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/

求根节点到叶节点数字之和

给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：

例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。

叶节点 是指没有子节点的节点。

提示：

树中节点的数目在范围 [1, 1000] 内
0 <= Node.val <= 9
树的深度不超过 10

*/

func main() {
	root := NewNormalTree()
	log.Println("求根节点到叶节点数字之和-dfs递归：", sumNumbers(root))
	log.Println("求根节点到叶节点数字之和-bfs迭代：", sumNumbers2(root))
}

func sumNumbers(root *TreeNode) int {
	var preorder func(node *TreeNode, sum int) int
	preorder = func(node *TreeNode, sum int) int {
		if node == nil {
			return 0
		}

		res := sum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return res
		}

		return preorder(node.Left, res) + preorder(node.Right, res)
	}

	return preorder(root, 0)
}

type pair struct {
	sum  int       // 累计值
	node *TreeNode // 当前结点
}

func sumNumbers2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	queue := make([]*pair, 0)
	queue = append(queue, &pair{sum: root.Val, node: root})
	for len(queue) != 0 {
		pop := queue[0]
		queue = queue[1:]

		if pop.node.Left == nil && pop.node.Right == nil {
			res += pop.sum
			continue
		}

		if pop.node.Left != nil {
			queue = append(queue, &pair{
				sum:  pop.sum*10 + pop.node.Left.Val,
				node: pop.node.Left,
			})
		}

		if pop.node.Right != nil {
			queue = append(queue, &pair{
				sum:  pop.sum*10 + pop.node.Right.Val,
				node: pop.node.Right,
			})
		}
	}

	return res
}
