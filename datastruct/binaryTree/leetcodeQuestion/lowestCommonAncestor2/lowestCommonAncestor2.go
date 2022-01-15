package main

import (
	"fmt"
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
二叉搜索树的最近公共祖先
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]

注意:1.最近公共祖先的val范围是[p, q] 2.不需要遍历，直接从根结点向下查找 3.p,q如果分布在root的两侧，那么就是root
4.在同一侧就一直查找

说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉搜索树中。

*/
func main() {
	root1 := NewSearchTreeNode()
	root2 := root1.Left.Left
	root3 := root1.Left.Right

	res := lowestCommonAncestor(root1, root2, root3)
	if res != nil {
		fmt.Println("搜索二叉树的【最近】公共祖先-递归：", res.Val)
	}
}

// lowestCommonAncestor
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// p,q其中之一是root，就是root
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	// p,q分布与root两侧，也是root
	if (root.Val > p.Val && root.Val < q.Val) || (root.Val < p.Val && root.Val > q.Val) {
		return root
	}

	intSize := 32 << (^uint(0) >> 63)
	max := -1 << (intSize - 1)
	min := -1 << (intSize - 1)
	if p.Val > q.Val {
		max = p.Val
		min = q.Val
	} else {
		max = q.Val
		min = p.Val
	}

	// p,q在root同一侧，只需要遍历一侧即可
	cur := root
	for cur != nil {
		if cur.Val >= min && cur.Val <= max {
			return cur
		}

		if cur.Val > max {
			cur = cur.Left
			continue
		}

		if cur.Val < min {
			cur = cur.Right
			continue
		}
	}

	return nil
}
