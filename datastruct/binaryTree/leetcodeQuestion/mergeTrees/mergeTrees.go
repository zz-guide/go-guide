package main

import (
	"fmt"
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/merge-two-binary-trees/
合并二叉树
给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为NULL 的节点将直接作为新二叉树的节点。

注意: 合并必须从两个树的根节点开始。

注意：
1.root1和root2同时存在左右结点的时候，需要新建结点
2.不同时存在，就把存在的树复用即可


*/
func main() {
	root1 := NewNormalTree()
	root2 := NewNormalTree2()

	log.Println("合并二叉树-递归：", levelorder.TraversalRecursive(mergeTrees(root1, root2)), levelorder.TraversalRecursive(root1), levelorder.TraversalRecursive(root2))
	log.Println("合并二叉树-迭代：", levelorder.TraversalRecursive(mergeTrees1(root1, root2)), levelorder.TraversalRecursive(root1), levelorder.TraversalRecursive(root2))
}

// mergeTrees, 递归,前序遍历
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  mergeTrees(root1.Left, root2.Left),
		Right: mergeTrees(root1.Right, root2.Right),
	}
}

// mergeTrees1 迭代方式，BFS
func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	var newRoot = &TreeNode{Val: root1.Val + root2.Val}

	var stack1 = []*TreeNode{newRoot} // 存放新建的结点，新建的结点需要链接子节点
	var stack2 = []*TreeNode{root1}
	var stack3 = []*TreeNode{root2}

	for len(stack2) > 0 && len(stack3) > 0 {
		// 出栈
		node1 := stack1[0]
		node2 := stack2[0]
		node3 := stack3[0]

		stack1 = stack1[1:]
		stack2 = stack2[1:]
		stack3 = stack3[1:]

		// 先判断左结点
		node2Left := node2.Left
		node3Left := node3.Left
		if node2Left != nil && node3Left != nil {
			tempNode := &TreeNode{Val: node2Left.Val + node3Left.Val}
			stack1 = append(stack1, tempNode)
			stack2 = append(stack2, node2Left)
			stack3 = append(stack3, node3Left)
			node1.Left = tempNode
		} else if node2Left != nil && node3Left == nil {
			node1.Left = node2Left
		} else if node2Left == nil && node3Left != nil {
			node1.Left = node3Left
		}

		// 然后是右节点
		node2Right := node2.Right
		node3Right := node3.Right
		if node2Right != nil && node3Right != nil {
			tempNode := &TreeNode{Val: node2Right.Val + node3Right.Val}
			stack1 = append(stack1, tempNode)
			stack2 = append(stack2, node2Right)
			stack3 = append(stack3, node3Right)
			node1.Right = tempNode
		} else if node2Right != nil && node3Right == nil {
			node1.Right = node2Right
		} else if node2Right == nil && node3Right != nil {
			node1.Right = node3Right
		}
	}

	return newRoot
}
