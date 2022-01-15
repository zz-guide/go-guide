package main

import (
	"fmt"
	. "go-guide/datastruct/binaryTree/treeNode"
)

/***
路径总和 II：https://leetcode-cn.com/problems/path-sum-ii/
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。
*/
func main() {
	root := NewNormalTree()

	targetSum := 11
	fmt.Println("所有路径总和II-递归：", pathSum(root, targetSum))
	//fmt.Println("所有路径总和II-递归：", pathSum1(root, targetSum))
	fmt.Println("所有路径总和II-递归：", pathSum2(root, targetSum))
}

// hasPathSum 递归法
func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var path []int

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}

		defer func() {
			path = path[:len(path)-1]
		}()

		left -= node.Val
		path = append(path, node.Val)

		if node.Left == nil && node.Right == nil && left == 0 {
			res = append(res, append([]int(nil), path...))
			return
		}

		dfs(node.Left, left)
		dfs(node.Right, left)
	}

	dfs(root, targetSum)
	return res
}

// pathSum1 迭代法BFS,前序遍历
/*func pathSum1(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	pathNodes := [][]*

	for len(queue) != 0 {
		n := queue[0]
		p := pathNodes[0]

		queue = queue[1:]
		pathNodes = pathNodes[1:]

		// 判断:当遍历到叶子节点时,如果p等于给定的值就返回真
		if n.Left == nil && n.Right == nil {
			if p == targetSum {
				return true
			}

			continue
		}

		if n.Left != nil {
			queue = append(queue, n.Left)
			pathSum = append(pathSum, n.Left.Val+p)
		}

		if n.Right != nil {
			queue = append(queue, n.Right)
			pathSum = append(pathSum, n.Right.Val+p)
		}
	}

	return res
}*/

type pair struct {
	node *TreeNode
	left int
}

// 层序遍历 bfs
func pathSum2(root *TreeNode, targetSum int) (ans [][]int) {
	if root == nil {
		return
	}

	parent := map[*TreeNode]*TreeNode{}

	getPath := func(node *TreeNode) (path []int) {
		for ; node != nil; node = parent[node] {
			path = append(path, node.Val)
		}
		for i, j := 0, len(path)-1; i < j; i++ {
			path[i], path[j] = path[j], path[i]
			j--
		}
		return
	}

	queue := []pair{{root, targetSum}}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		node := p.node
		left := p.left - node.Val
		if node.Left == nil && node.Right == nil {
			if left == 0 {
				ans = append(ans, getPath(node))
			}
		} else {
			if node.Left != nil {
				parent[node.Left] = node
				queue = append(queue, pair{node.Left, left})
			}
			if node.Right != nil {
				parent[node.Right] = node
				queue = append(queue, pair{node.Right, left})
			}
		}
	}

	return
}
