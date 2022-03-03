package main

import (
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

从前序与中序遍历序列构造二叉树

提示:

1 <= preorder.length <= 3000
inorder.length == preorder.length
-3000 <= preorder[i], inorder[i] <= 3000
preorder和inorder均无重复元素
inorder均出现在preorder
preorder保证为二叉树的前序遍历序列
inorder保证为二叉树的中序遍历序列


*/
func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	log.Println("从前序与中序遍历序列构造二叉树-递归：", levelorder.TraversalRecursive(buildTree(preorder, inorder)))
	log.Println("从前序与中序遍历序列构造二叉树-迭代法：", levelorder.TraversalRecursive(buildTree1(preorder, inorder)))
}

// buildTree 递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]

	var rootIndex int
	for rootIndex = range inorder {
		if inorder[rootIndex] == rootVal {
			break
		}
	}

	// 前序：中左右
	// 中序：左中右
	// 技巧就是想办法找到左子树和右子树的边界，分别传入就可以了

	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(preorder[1:len(inorder[:rootIndex])+1], inorder[:rootIndex]),
		Right: buildTree(preorder[len(inorder[:rootIndex])+1:], inorder[rootIndex+1:]),
	}
}

// buildTree1 迭代法
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	var stack []*TreeNode
	stack = append(stack, root)
	var inorderIndex int
	for i := 1; i < len(preorder); i++ {
		// 前序遍历，最前边的值都是根节点
		node := stack[len(stack)-1]
		// 一直向左移动，直到最左
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{Val: preorder[i]}
			stack = append(stack, node.Left)
		} else {
			// 此处不能直接使用node，因为node会不断回溯向上寻找
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{Val: preorder[i]}
			stack = append(stack, node.Right)
		}
	}
	return root
}
