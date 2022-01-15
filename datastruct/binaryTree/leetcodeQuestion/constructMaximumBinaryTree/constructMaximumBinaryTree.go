package main

import (
	"fmt"
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/maximum-binary-tree/

最大二叉树

给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：

二叉树的根是数组 nums 中的最大元素。
左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
返回有给定数组 nums 构建的 最大二叉树 。

提示：

1 <= nums.length <= 1000
0 <= nums[i] <= 1000
nums 中的所有整数 互不相同


注意：1.遍历数组从左往右
2.左节点只能从左侧找
3.右结点只能从右侧找

*/
func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	fmt.Println("最大二叉树-递归：", levelorder.TraversalRecursive(constructMaximumBinaryTree(nums)))
	fmt.Println("最大二叉树-迭代：", levelorder.TraversalRecursive(constructMaximumBinaryTree1(nums)))
}

// constructMaximumBinaryTree 递归法
func constructMaximumBinaryTree(nums []int) *TreeNode {
	var helper func(nums []int) *TreeNode
	helper = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}

		maxIndex := 0
		for i, num := range nums {
			if num > nums[maxIndex] {
				maxIndex = i
			}
		}

		node := &TreeNode{Val: nums[maxIndex]}
		node.Left = helper(nums[:maxIndex])
		node.Right = helper(nums[maxIndex+1:])

		return node
	}

	return helper(nums)
}

// constructMaximumBinaryTree1 迭代法
func constructMaximumBinaryTree1(nums []int) *TreeNode {
	stack := make([]*TreeNode, 0, len(nums))

	// 因为遍历数组是从左往右的，第一个数只可能有右节点，肯定没有左节点
	for _, num := range nums {
		curr := &TreeNode{Val: num}
		size := len(stack)
		for size > 0 && stack[size-1].Val < num {
			// 当前结点大，只能在当前结点的左侧
			curr.Left = stack[size-1]
			stack = stack[:size-1]
			size = len(stack)
		}

		if len(stack) > 0 {
			stack[len(stack)-1].Right = curr
		}

		// 依次入栈
		stack = append(stack, curr)
	}

	return stack[0]
}
