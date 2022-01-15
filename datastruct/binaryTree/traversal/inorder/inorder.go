package inorder

import (
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
结论：中序遍历(升序遍历)，遍历顺序：左子树->根->右子树
*/

// TraversalRecursive 函数递归遍历
func TraversalRecursive(root *TreeNode) []int {
	var res []int
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		preorder(node.Left)
		res = append(res, node.Val)
		preorder(node.Right)
	}

	preorder(root)
	return res
}

// TraversalRecursive1 采用递归的方式, 不使用函数内闭包的形式，需要2个方法
func TraversalRecursive1(root *TreeNode) []int {
	return inorderRecursive(root)
}

func inorderRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	rest := append(inorderRecursive(root.Left), root.Val)
	rest = append(rest, inorderRecursive(root.Right)...)
	return rest
}

// TraversalStack 使用栈辅助，比函数递归要更省空间一点
func TraversalStack(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	node := root

	for node != nil || 0 < len(stack) {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 中序
		res = append(res, node.Val)
		node = node.Right
	}

	return res
}

// TraversalStack1 标记法，本质也是栈，缺点是每个结点都出入两次栈，总体效率要比递归好
func TraversalStack1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := []MarkNode{{IsVisit: false, Node: root}}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !cur.IsVisit {
			if cur.Node.Right != nil {
				stack = append(stack, MarkNode{IsVisit: false, Node: cur.Node.Right})
			}

			stack = append(stack, MarkNode{IsVisit: true, Node: cur.Node})

			if cur.Node.Left != nil {
				stack = append(stack, MarkNode{IsVisit: false, Node: cur.Node.Left})
			}

		} else {
			res = append(res, cur.Node.Val)
		}
	}

	return res
}
