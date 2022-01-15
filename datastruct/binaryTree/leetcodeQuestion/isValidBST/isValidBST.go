package main

import (
	"fmt"
	"math"

	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
题目：https://leetcode-cn.com/problems/validate-binary-search-tree/submissions/
验证二叉搜索树

给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

注意：本质上是利用中序遍历的结果是升序来判断
*/
func main() {
	searchNode := NewSearchTreeNode()
	fmt.Println("验证二叉搜索树-迭代法:", isValidBST(searchNode))
	fmt.Println("验证二叉搜索树-Morris遍历判断:", TraversalMorris(searchNode))
}

func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

func TraversalMorris(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var morrisRes []int  // 存放cur每次遍历ur结点的值
	var inorderRes []int // 存放中序遍历

	var mostRight *TreeNode // 存放左子树的最右子树结点
	cur := root
	var pre *TreeNode

	// 等同于while循环，cur为nil退出循环，也就遍历完了
	for cur != nil {

		morrisRes = append(morrisRes, cur.Val)

		// 首先把cur的左子树赋值给mostRight
		mostRight = cur.Left
		if mostRight != nil { // 存在左子树的逻辑
			// 不停的寻找最右子树，mostRight.Right != cur用来判断是首次遍历
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}

			// 首次遍历到这个最右结点，更改指向，并向左移动
			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else { // 再次遍历到最右结点，恢复nil指向，向右移动
				if pre != nil && pre.Val >= cur.Val {
					return false
				}

				pre = cur

				// 中序遍历
				inorderRes = append(inorderRes, cur.Val)

				mostRight.Right = nil

				// 这2行其实可以不要，因为 最外层else也是右移
				cur = cur.Right
				continue
			}

		} else {
			if pre != nil && pre.Val >= cur.Val {
				return false
			}

			pre = cur
			// 收集中序遍历的值
			inorderRes = append(inorderRes, cur.Val)
			cur = cur.Right // 不存在左子树就右移
		}
	}

	return true
}
