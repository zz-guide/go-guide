package postorder

import (
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
结论：后序遍历类似归并排序的思想，遍历顺序：左子树->右子树->根
*/

func TraversalRecursive(root *TreeNode) []int {
	var res []int
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return res
}

// TraversalRecursive1 采用递归的方式, 不使用函数内闭包的形式，需要2个方法
func TraversalRecursive1(root *TreeNode) []int {
	return postorderRecursive(root)
}

func postorderRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	// 因为。。。不能是首位，所以需要2个变量
	l := append(postorderRecursive(root.Left))
	r := append(postorderRecursive(root.Right), root.Val)

	//return append(
	//	postorderRecursive(root.Left),
	//	append(postorderRecursive(root.Right),root.Val)...)
	return append(l, r...)
}

// TraversalStack 使用栈辅助，比函数递归要更省空间一点
func TraversalStack(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	var pre *TreeNode

	node := root
	for node != nil || 0 < len(stack) {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Right == nil || node.Right == pre {
			res = append(res, node.Val)
			// 防止右结点回到根节点以后又向下遍历了
			pre = node
			node = nil
		} else {
			// 重新压栈，之后能顺着往上层遍历
			stack = append(stack, node)
			node = node.Right
		}
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
			stack = append(stack, MarkNode{IsVisit: true, Node: cur.Node})

			if cur.Node.Right != nil {
				stack = append(stack, MarkNode{IsVisit: false, Node: cur.Node.Right})
			}

			if cur.Node.Left != nil {
				stack = append(stack, MarkNode{IsVisit: false, Node: cur.Node.Left})
			}

		} else {
			// 访问过直接存结果
			res = append(res, cur.Node.Val)
		}
	}

	return res
}

// TraversalStack2 后序遍历=前序遍历左右颠倒，再取反
func TraversalStack2(root *TreeNode) []int {
	var f func(root *TreeNode) []int
	f = func(root *TreeNode) []int {
		var res []int
		var stack []*TreeNode
		node := root

		// 以中右左的形式输出
		for node != nil || 0 < len(stack) {
			for node != nil {
				res = append(res, node.Val)
				stack = append(stack, node)
				node = node.Right
			}

			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node = node.Left
		}

		return res
	}

	// 左和右互换位置就是最后输出结果
	res := f(root)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	return res
}
