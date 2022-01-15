package main

import (
	"fmt"
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/search-in-a-binary-search-tree/
二叉搜索树中的搜索

给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。

例如，

给定二叉搜索树:

        4
       / \
      2   7
     / \
    1   3

和值: 2
你应该返回如下子树:

      2
     / \
    1   3
在上述示例中，如果要找的值是 5，但因为没有节点值为 5，我们应该返回 NULL。

*/
func main() {
	root := NewSearchTreeNode()
	val := 2
	fmt.Println("二叉搜索树中的搜索-递归：", levelorder.TraversalRecursive(searchBST(root, val)))
	fmt.Println("二叉搜索树中的搜索-迭代：", levelorder.TraversalRecursive(searchBST1(root, val)))
}

// 递归法，适用于所有的二叉树
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val == root.Val {
		return root
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	}

	if val > root.Val {
		return searchBST(root.Right, val)
	}

	return nil
}

// 利用搜索二叉树的性质，不递归可以快速查找
func searchBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	cur := root
	for cur != nil {
		if cur.Val == val {
			return cur
		}

		if cur.Val > val {
			cur = cur.Left
			continue
		}

		if cur.Val < val {
			cur = cur.Right
			continue
		}
	}

	return nil
}
