package main

import (
	"fmt"
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
二叉树的最近公共祖先
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

提示：

树中节点数目在范围 [2, 105] 内。
-109 <= Node.val <= 109
所有 Node.val 互不相同。
p != q
p 和 q 均存在于给定的二叉树中


注意:
1.如果root是p,q中任意一个，这时候的公共祖先就是root
2.root的左子树有p或者q,右子树有p或者q,root是公共祖先
3.p,q同时在左子树或者右子树，需要递归root子树求解

*/
func main() {
	root1 := NewNormalTree()
	root2 := root1.Left.Left
	root3 := root1.Left.Right

	res := lowestCommonAncestor(root1, root2, root3)
	if res != nil {
		fmt.Println("二叉树的【最近】公共祖先-递归：", res.Val)
	}
}

// lowestCommonAncestor 后序遍历，递归，dfs
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 条件1
	if root == p || root == q {
		return root
	}

	// 函数返回当前结点左子树的最近祖先结点
	leftAncestor := lowestCommonAncestor(root.Left, p, q)
	rightAncestor := lowestCommonAncestor(root.Right, p, q)

	// 条件2
	if leftAncestor != nil && rightAncestor != nil {
		return root
	}

	if leftAncestor != nil {
		return leftAncestor
	}

	if rightAncestor != nil {
		return rightAncestor
	}

	return nil
}

// lowestCommonAncestor1 先递归寻找q的祖先结点，然后遍历p,最先碰到的就是最近祖先
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			parent[r.Left.Val] = r
			dfs(r.Left)
		}
		if r.Right != nil {
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)

	ancestors := map[int]bool{}
	for p != nil {
		ancestors[p.Val] = true
		p = parent[p.Val]
	}

	for q != nil {
		// 如果发现已经设置为祖先了，那么就是最近的
		if ancestors[q.Val] {
			return q
		}

		// 不是就向父级方向移动
		q = parent[q.Val]
	}

	return nil
}
