package treeNode

import "math/rand"

// TreeNode 二叉树结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

// NewTreeNode 随意构造一颗二叉树
func NewTreeNode(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	// 此处的mid可以是中间数的左边或者右边，此处随机取
	mid := (left + right + rand.Intn(2)) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}

// NewSearchTreeNode 构造一颗二叉搜索树，要求nums升序
func NewSearchTreeNode() *TreeNode {
	/**
			 4
		2          6
	1      3    5     7
	*/
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	root.Right = &TreeNode{Val: 6}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func NewNormalTree() *TreeNode {
	/**
			 1
		2          3
	4      5    6     7
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func NewNormalTree2() *TreeNode {
	/**
			 2
		1           3
	       4    6      7
	*/
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 4}

	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// NewSymmetricTree 构造一颗对称二叉树
func NewSymmetricTree() *TreeNode {
	/**
			 1
		2          2
	3      4    4     3
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}

	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}
	return root
}

// NewCompleteBinaryTree 构造一颗完全二叉树
func NewCompleteBinaryTree() *TreeNode {
	/**
			 1
		2          3
	4      5    6
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	return root
}

// NewBalanceBinaryTree 构造一颗平衡二叉树
func NewBalanceBinaryTree() *TreeNode {
	/**
			   3
		9           20
	            15      7
	*/
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}

	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func NewInsertSearchTreeNode() *TreeNode {
	/**
			 7
		2           11
	1      3    10      12
	*/
	root := &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	root.Right = &TreeNode{Val: 11}
	root.Right.Left = &TreeNode{Val: 10}
	root.Right.Right = &TreeNode{Val: 12}
	return root
}
