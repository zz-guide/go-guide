package main

import (
	"fmt"
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/house-robber-iii/
打家劫舍 III

在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。
这个地区只有一个入口，我们称之为“根”。
除了“根”之外，每栋房子有且只有一个“父“房子与之相连。
一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。

计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。


注意：1.假设偷当前结点，就不能偷自己的左右孩子,可以偷孩子的孩子(如果有的话)
2.假设不偷当前结点，可以偷左右孩子
3.树形dp其实就是递归

*/
func main() {
	nums := NewNormalTree()
	fmt.Println("打家劫舍III-递归:", rob(nums))
	fmt.Println("打家劫舍III-记忆化递归:", rob1(nums))
	fmt.Println("打家劫舍III-动态规划+递归:", rob2(nums))
}

// rob 递归, 会超时
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 偷根节点
	robRoot := root.Val
	if root.Left != nil {
		robRoot += rob(root.Left.Left) + rob(root.Left.Right)
	}

	if root.Right != nil {
		robRoot += rob(root.Right.Left) + rob(root.Right.Right)
	}

	// 不偷根节点
	robNoRoot := rob(root.Left) + rob(root.Right)
	if robRoot > robNoRoot {
		return robRoot
	}

	return robNoRoot
}

// rob1 记忆化递归
func rob1(root *TreeNode) int {

	memo := map[*TreeNode]int{}

	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		if _, ok := memo[root]; ok {
			return memo[root]
		}

		robIncludeRoot := root.Val
		if root.Left != nil {
			robIncludeRoot += rob(root.Left.Left) + rob(root.Left.Right)
		}

		if root.Right != nil {
			robIncludeRoot += rob(root.Right.Left) + rob(root.Right.Right)
		}

		robExcludeRoot := rob(root.Left) + rob(root.Right)

		res := 0
		if robIncludeRoot > robExcludeRoot {
			res = robIncludeRoot
		} else {
			res = robExcludeRoot
		}

		memo[root] = res
		return res
	}

	return helper(root)
}

func rob2(root *TreeNode) int {
	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			// 第一个值代表偷的最大金额，第二个值代表不偷的最大金额
			return []int{0, 0}
		}

		l, r := dfs(node.Left), dfs(node.Right)
		selected := node.Val + l[1] + r[1]
		notSelected := max(l[0], l[1]) + max(r[0], r[1])
		return []int{selected, notSelected}
	}

	val := dfs(root)
	return max(val[0], val[1])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
