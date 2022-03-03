package main

import (
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/recover-binary-search-tree/

恢复二叉搜索树

给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。


示例 1：


输入：root = [1,3,null,null,2]
输出：[3,1,null,null,2]
解释：3 不能是 1 的左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。
示例 2：


输入：root = [3,1,4,null,null,2]
输出：[2,1,4,null,null,3]
解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。

提示：

树上节点的数目在范围 [2, 1000] 内
-231 <= Node.val <= 231 - 1

进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用 O(1) 空间的解决方案吗？

*/
func main() {
	root1 := NewWrongSearchTree()
	root2 := NewWrongSearchTree()
	recoverTree(root1)
	log.Println("恢复二叉搜索树-迭代：", levelorder.TraversalRecursive(root1))
	recoverTree2(root2)
	log.Println("恢复二叉搜索树-Morris遍历：", levelorder.TraversalRecursive(root2))
}

// recoverTree 迭代
func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	var stack []*TreeNode
	var first, second, preNode *TreeNode // first,y就是这两个顺序不对的结点,preNode表示迭代过程中的上一个结点值
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 如果当前结点的值小于preNode，说明不符合遍历顺序
		// 两个结点有可能挨着，有可能隔开，所以需要交替值
		if preNode != nil && root.Val < preNode.Val {
			// 先赋值second
			second = root
			if first == nil {
				first = preNode
			} else {
				break
			}
		}

		// 更新preNode
		preNode = root

		root = root.Right
	}
	first.Val, second.Val = second.Val, first.Val
}

// recoverTree2 Morris遍历
func recoverTree2(root *TreeNode) {
	if root == nil {
		return
	}

	var first, second, preNode, predecessor *TreeNode

	for root != nil {
		if root.Left != nil {
			predecessor = root.Left
			// 一直找左子树的右孩子，但是不能形成环
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}

			// 让 predecessor 的右指针指向 root，继续遍历左子树
			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
			} else { // 说明左子树已经访问完了，我们需要断开链接
				// 这段逻辑是公用的
				{
					if preNode != nil && root.Val < preNode.Val {
						second = root
						if first == nil {
							first = preNode
						}
					}
					preNode = root
				}

				// 解环，向右移动
				predecessor.Right = nil
				root = root.Right
			}
		} else { // 如果没有左孩子，则直接访问右孩子
			// 这段逻辑是公用的，相比原来的morris遍历就多了这么一段
			{
				if preNode != nil && root.Val < preNode.Val {
					second = root
					if first == nil {
						first = preNode
					}
				}
				preNode = root
			}

			root = root.Right
		}
	}
	// 交换结点的值即可，因为要保持结构不变
	first.Val, second.Val = second.Val, first.Val
}
