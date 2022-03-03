package main

import (
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/balance-a-binary-search-tree/submissions/
将二叉搜索树变平衡
高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。


给你一棵二叉搜索树，请你返回一棵平衡后的二叉搜索树，新生成的树应该与原来的树有着相同的节点值。

如果一棵二叉搜索树中，每个节点的两棵子树高度差不超过 1 ，我们就称这棵二叉搜索树是 平衡的 。

如果有多种构造方法，请你返回任意一种。

注意：1.中序遍历正好是升序数组
2.升序+高度平衡无法确定唯一二叉树，根节点不固定
3.通过数组的二分查找算法可知，二分后的区间构成的二叉树就是高度平衡的
4.还有一种方法是AVL旋转，但是实现很麻烦，还得旋转结点，效率不高
*/
func main() {
	root := NewSearchTreeNode()

	log.Println("将二叉搜索树变平衡-递归：", levelorder.TraversalRecursive(balanceBST1(root)))
}

// balanceBST 二叉搜索树可以通过中序遍历得到升序数组，然后二分构造平衡树,构造出来的树答案不唯一
func balanceBST1(root *TreeNode) *TreeNode {
	// 先中序遍历
	var inorder func(root *TreeNode) []int
	inorder = func(root *TreeNode) []int {
		var res []int
		var helper func(root *TreeNode)
		helper = func(root *TreeNode) {
			if root == nil {
				return
			}

			helper(root.Left)
			res = append(res, root.Val)
			helper(root.Right)
		}
		helper(root)
		return res
	}

	// 二分构造树
	var buildTree func(nums []int) *TreeNode
	buildTree = func(nums []int) *TreeNode {
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

	return buildTree(inorder(root))
}
