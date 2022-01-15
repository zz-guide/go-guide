package main

import (
	"fmt"
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
	"math/rand"
)

/**
题目：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。

高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

注意：1.中序遍历正好是升序数组
2.升序+高度平衡无法确定唯一二叉树，根节点不固定
3.通过数组的二分查找算法可知，二分后的区间构成的二叉树就是高度平衡的
*/
func main() {
	nums := []int{-10, -3, 0, 5, 9}

	fmt.Println("构建平衡二叉树-递归-中左：", levelorder.TraversalRecursive(sortedArrayToBST(nums)))
	fmt.Println("构建平衡二叉树-递归-中右：", levelorder.TraversalRecursive(sortedArrayToBST1(nums)))
	fmt.Println("构建平衡二叉树-递归-中随意：", levelorder.TraversalRecursive(sortedArrayToBST2(nums)))
}

// sortedArrayToBST 中序遍历，总是选择中间位置左边的数字作为根节点
func sortedArrayToBST(nums []int) *TreeNode {
	var helper func(left, right int) *TreeNode
	helper = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (right-left)/2 + left
		root := &TreeNode{Val: nums[mid]}
		root.Left = helper(left, mid-1)
		root.Right = helper(mid+1, right)
		return root
	}

	return helper(0, len(nums)-1)
}

// sortedArrayToBST1 中序遍历，总是选择中间位置右边的数字作为根节点
func sortedArrayToBST1(nums []int) *TreeNode {
	var helper func(left, right int) *TreeNode
	helper = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		// 总是选择中间位置右边的数字作为根节点
		mid := (left + right + 1) / 2
		root := &TreeNode{Val: nums[mid]}
		root.Left = helper(left, mid-1)
		root.Right = helper(mid+1, right)
		return root
	}

	return helper(0, len(nums)-1)
}

// sortedArrayToBST2 中序遍历，选择任意一个中间位置数字作为根节点
func sortedArrayToBST2(nums []int) *TreeNode {
	var helper func(left, right int) *TreeNode
	helper = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		// 选择任意一个中间位置数字作为根节点
		mid := (left + right + rand.Intn(2)) / 2
		root := &TreeNode{Val: nums[mid]}
		root.Left = helper(left, mid-1)
		root.Right = helper(mid+1, right)
		return root
	}

	return helper(0, len(nums)-1)
}
