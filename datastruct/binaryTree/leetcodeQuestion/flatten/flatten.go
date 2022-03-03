package main

import (
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/

二叉树展开为链表

给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [0]
输出：[0]

提示：

树中结点数在范围 [0, 2000] 内
-100 <= Node.val <= 100

进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？

*/
func main() {
	root := NewSearchTreeNode()
	flatten(root)
	log.Println("将二叉搜索树变平衡-递归：", levelorder.TraversalRecursive(root))
}

// flatten 递归前序遍历
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	pre := &TreeNode{}
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		left, right := node.Left, node.Right
		// 当前结点左右结点置位nil
		node.Left, node.Right = nil, nil
		// 上一个结点right指向当前node
		pre.Right = node
		// 移动
		pre = pre.Right

		preorder(left)
		preorder(right)
	}

	preorder(root)
	return
}

// flatten2 Morris前序遍历
func flatten1(root *TreeNode) {
	cur := root
	for cur != nil {
		left, right := cur.Left, cur.Right
		predecessor := left
		if left != nil {
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = right
			cur.Left = nil
			cur.Right = left
		}
		cur = cur.Right
	}
}
