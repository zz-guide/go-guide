package preorder

import (
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
结论：前序遍历类似快速排序的思想，遍历顺序：根->左子树->右子树
*/

// TraversalRecursive 函数递归遍历
func TraversalRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}

	preorder(root)
	return res
}

// TraversalRecursive1 采用递归的方式, 不使用函数内闭包的形式，需要2个方法
func TraversalRecursive1(root *TreeNode) []int {
	return preorderRecursive(root)
}

func preorderRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	rest := append([]int{root.Val}, preorderRecursive(root.Left)...)
	rest = append(rest, preorderRecursive(root.Right)...)
	return rest
}

// TraversalStack 使用栈辅助，比函数递归要更省空间一点
func TraversalStack(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var stack []*TreeNode
	node := root

	// 对于root结点。先输出，向左移动，最后一步步从底向上回溯，然后遍历右边
	for node != nil || len(stack) > 0 {
		for node != nil {
			// 前序
			res = append(res, node.Val)
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
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
		// 从尾部弹出，注意append的顺序
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !cur.IsVisit {

			if cur.Node.Right != nil {
				stack = append(stack, MarkNode{IsVisit: false, Node: cur.Node.Right})
			}

			if cur.Node.Left != nil {
				stack = append(stack, MarkNode{IsVisit: false, Node: cur.Node.Left})
			}

			stack = append(stack, MarkNode{IsVisit: true, Node: cur.Node})

		} else {
			// 访问过直接存结果
			res = append(res, cur.Node.Val)
		}
	}

	return res
}
