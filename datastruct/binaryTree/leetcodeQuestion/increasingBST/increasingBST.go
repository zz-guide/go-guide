package main

import (
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/NYBBNL/solution/zhan-ping-er-cha-sou-suo-shu-by-leetcode-pmxr/
展平二叉搜索树

给你一棵二叉搜索树，请 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右子节点。
提示：

树中节点数的取值范围是 [1, 100]
0 <= Node.val <= 1000

注意：
1.二叉搜索树
2.最左边的结点要成为根节点，也就是最小值是跟结点
3.新的二叉搜索树只有右节点，没有左节点
4.可以直接修改指向，按照前序遍历的顺序指即可

*/
func main() {
	root1 := NewSearchTreeNode()
	root2 := NewSearchTreeNode()

	log.Println("二叉树的镜像-直接遍历：", levelorder.TraversalRecursive(increasingBST(root1)))
	log.Println("二叉树的镜像-修改指向：", levelorder.TraversalRecursive(increasingBST1(root2)))
}

// increasingBST 前序遍历，先遍历，后建立
func increasingBST(root *TreeNode) *TreeNode {
	var res []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)
			res = append(res, node.Val)
			inorder(node.Right)
		}
	}

	inorder(root)

	dummyNode := &TreeNode{}
	curNode := dummyNode
	for _, val := range res {
		curNode.Right = &TreeNode{Val: val}
		curNode = curNode.Right
	}
	return dummyNode.Right
}

// increasingBST1 前序遍历，直接修改指向
func increasingBST1(root *TreeNode) *TreeNode {
	dummyNode := &TreeNode{}
	resNode := dummyNode

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		// 在中序遍历的过程中修改节点指向
		resNode.Right = node
		node.Left = nil
		resNode = node

		inorder(node.Right)
	}
	inorder(root)

	return dummyNode.Right
}
