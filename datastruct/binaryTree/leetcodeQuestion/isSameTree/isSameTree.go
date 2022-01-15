package main

import (
	"fmt"
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/same-tree/
相同的树

给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。


注意：1.四种遍历都可以的 2.逐个比较结点的值和位置即可

*/
func main() {
	root1 := NewSearchTreeNode()
	root2 := NewSearchTreeNode()

	fmt.Println("相同的树-递归：", isSameTree(root1, root2))
	fmt.Println("相同的树-迭代：", isSameTree1(root1, root2))
}

// isSameTree 递归
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 层序遍历
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	levelNodeQueue := []*TreeNode{p, q}
	// 以p为参照物
	for len(levelNodeQueue) > 0 {
		length := len(levelNodeQueue)
		for j := 0; j < length; j = j + 2 {

			pNode := levelNodeQueue[0]
			qNode := levelNodeQueue[1]
			levelNodeQueue = levelNodeQueue[2:]
			if pNode == nil && qNode == nil {
				continue
			}

			if (pNode == nil && qNode != nil) || (pNode != nil && qNode == nil) {
				return false
			}
			if pNode != nil && qNode != nil && pNode.Val != qNode.Val {
				return false
			}

			if pNode != nil {
				levelNodeQueue = append(levelNodeQueue, pNode.Left)
			} else {
				levelNodeQueue = append(levelNodeQueue, nil)
			}

			if qNode != nil {
				levelNodeQueue = append(levelNodeQueue, qNode.Left)
			} else {
				levelNodeQueue = append(levelNodeQueue, nil)
			}

			if pNode != nil {
				levelNodeQueue = append(levelNodeQueue, pNode.Right)
			} else {
				levelNodeQueue = append(levelNodeQueue, nil)
			}

			if qNode != nil {
				levelNodeQueue = append(levelNodeQueue, qNode.Right)
			} else {
				levelNodeQueue = append(levelNodeQueue, nil)
			}

		}

	}

	return true
}
