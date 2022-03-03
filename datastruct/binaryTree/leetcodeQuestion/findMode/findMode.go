package main

import (
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/
二叉搜索树中的众数
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
例如：
给定 BST [1,null,2,2],

   1
    \
     2
    /
   2
返回[2].

提示：如果众数超过1个，不需考虑输出顺序

进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）





*/
func main() {
	root := NewSearchTreeNode()
	log.Println("二叉搜索树中的众数-递归：", findMode(root))
	log.Println("二叉搜索树中的众数-迭代：", findMode(root))
}

// findMode 中序遍历，递归法
func findMode(root *TreeNode) []int {
	var res []int
	var count int // 计数器
	var maxFrequency int
	var preNode *TreeNode

	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}

		helper(node.Left)

		if preNode != nil && preNode.Val == node.Val {
			count++
		} else {
			count = 1
		}

		if count >= maxFrequency {
			// 大于就覆盖，等于就添加
			if count > maxFrequency {
				res = []int{node.Val}
				maxFrequency = count
			} else if count == maxFrequency {
				res = append(res, node.Val)
			}
		}

		preNode = node
		helper(node.Right)
	}

	helper(root)
	return res
}

// findMode1 Morris遍历
func findMode1(root *TreeNode) []int {
	var base, count, maxCount int
	var res []int
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			res = append(res, base)
		} else if count > maxCount {
			maxCount = count
			res = []int{base}
		}
	}

	cur := root
	for cur != nil {
		if cur.Left == nil {
			update(cur.Val)
			cur = cur.Right
			continue
		}

		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}

		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
		} else {
			pre.Right = nil
			update(cur.Val)
			cur = cur.Right
		}
	}

	return res
}

// findMode2 迭代法
func findMode2(root *TreeNode) []int {
	var base, count, maxCount int
	var res []int
	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			res = append(res, base)
		} else if count > maxCount {
			maxCount = count
			res = []int{base}
		}
	}

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
		update(node.Val)
		node = node.Right
	}

	return res
}
