package main

import (
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

根据一棵树的中序遍历与后序遍历构造二叉树。

注意:你可以假设树中没有重复的元素。  这个条件很重要，不然不好在中序数组中找跟元素

例如，给出

中序遍历 inorder =[9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7


*/
func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	log.Println("中序后序构造二叉树-递归：", levelorder.TraversalRecursive(buildTree(inorder, postorder)))
	log.Println("中序后序构造二叉树-迭代：", levelorder.TraversalRecursive(buildTree1(inorder, postorder)))
}

// buildTree 递归
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1] // 后序数组的最后一个元素就是根元素

	var rootIndex int
	for rootIndex = range inorder {
		if inorder[rootIndex] == rootVal {
			break
		}
	}

	// 中序：左中右
	// 后序：左右中
	// 只是根和右子树顺序不同，对于每一个根节点来说，中序根节点的位置正好是后序遍历左右子树分开的位置，不包含最后一个元素
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(inorder[:rootIndex], postorder[:rootIndex]),
		Right: buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1]),
	}
}

// buildTree1 迭代法
func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	stack := []*TreeNode{root}
	inorderIndex := len(inorder) - 1

	// 中序：左中右
	// 后序：左右中
	// i 从根节点开始构建
	for i := len(postorder) - 2; i >= 0; i-- {
		// 当前结点的值
		postorderVal := postorder[i]
		node := stack[len(stack)-1]
		// 先找最右结点
		if node.Val != inorder[inorderIndex] {
			node.Right = &TreeNode{Val: postorderVal}
			stack = append(stack, node.Right)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex--
			}
			node.Left = &TreeNode{Val: postorderVal}
			stack = append(stack, node.Left)
		}
	}
	return root
}
