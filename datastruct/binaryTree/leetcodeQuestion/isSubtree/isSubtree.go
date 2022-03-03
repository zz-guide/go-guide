package main

import (
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/subtree-of-another-tree/
另一棵树的子树

给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。如果存在，返回 true ；否则，返回 false 。

二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。

题解：1.递归判断root的每一颗子树是不是和subRoot相等或者对称
2.将二叉树串成一个字符串，然后搜索字符串，可以通过KMP算法
3.还可以使用hash思想
*/
func main() {
	root1 := NewSearchTreeNode()
	root2 := NewSearchTreeNode()

	log.Println("另一颗树的子树-递归：", isSubtree(root1, root2))
	log.Println("另一颗树的子树-迭代：", isSubtree1(root1, root2))
}

// 用相等或者对称的逻辑都可以
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

// isSameTree 递归
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// isSubtree1 层序遍历
func isSubtree1(p *TreeNode, q *TreeNode) bool {
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
