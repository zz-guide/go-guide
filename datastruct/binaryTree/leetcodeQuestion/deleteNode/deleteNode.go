package main

import (
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/delete-node-in-a-bst/
删除二叉搜索树中的节点

给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。

提示:

节点数的范围[0, 104].
-105<= Node.val <= 105
节点值唯一
root是合法的二叉搜索树
-105<= key <= 105



题解：
1.如果被删除节点是 leaf, 直接删除
2.如果被删除节点 只有一个child, 使用仅有的 child 代替原节点
3.如果被删除节点 有两个children, 则在 right subtree中 寻找 successor, 将原节点值替换为 successor 的值, 并递归删除 successor, 将问题转化为 情况1 或 情况2.


*/
func main() {
	root := NewSearchTreeNode()
	val := 2
	log.Println("删除二叉搜索树结点-递归：", levelorder.TraversalRecursive(deleteNode(root, val)))
	log.Println("删除二叉搜索树结点-迭代：", levelorder.TraversalRecursive(deleteNode1(root, val)))
}

// deleteNode 递归法
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// 查找删除的值在哪一侧
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		// 如果当前结点正好是要删除的key

		// 左子树是空的，返回右子树
		if root.Left == nil {
			return root.Right
		}

		// 右子树是空的，返回左子树
		if root.Right == nil {
			return root.Left
		}

		// 都不为空，左子树需要下放到右子树的最左侧
		temp := root.Right
		for temp.Left != nil {
			temp = temp.Left
		}

		temp.Left = root.Left
		root = root.Right
	}

	return root
}

// deleteNode1 迭代版本
func deleteNode1(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// 1.二叉搜索树dfs查找到要删除的结点，并且记录是左树还是右树，然后是父节点
	delNode := root
	var parent *TreeNode
	isLeft, isRight := false, false
	for delNode != nil {
		if delNode.Val == key {
			break
		}

		parent = delNode
		if delNode.Val > key {
			delNode = delNode.Left
			isLeft = true
			isRight = false
		} else {
			delNode = delNode.Right
			isRight = true
			isLeft = false
		}
	}

	// 2.没找到要删除的结点
	if delNode == nil {
		return root
	}

	// 3.修剪要删除的这课子树，同时计算最终要挂在父节点的这个结点
	var tempNode *TreeNode
	if delNode.Left == nil && delNode.Right == nil {
		tempNode = nil
	} else if delNode.Left == nil && delNode.Right != nil {
		tempNode = delNode.Right
	} else if delNode.Right == nil && delNode.Left != nil {
		tempNode = delNode.Left
	} else {
		tempRight := delNode.Right
		for tempRight.Left != nil {
			tempRight = tempRight.Left
		}

		tempRight.Left = delNode.Left
		tempNode = delNode.Right
	}

	// 4.父节点可能没有，比如root这个结点就是要删除的结点的时候
	if parent == nil {
		return tempNode
	}

	// 5.挂载到父结点
	if isLeft {
		parent.Left = tempNode
	} else if isRight {
		parent.Right = tempNode
	}

	return root
}
