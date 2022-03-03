package main

import (
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/
填充每个节点的下一个右侧节点指针

给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有next 指针都被设置为 NULL。


进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

注意：题目要求把每一层结点用另外一个指针链接起来，其实就是层序遍历

*/
func main() {
	root := NewNormalTree()

	log.Println("填充每一层结点右侧指针-迭代O(N)：", connect(root))
	log.Println("填充每一层结点右侧指针-递归", connect2(root))
	log.Println("填充每一层结点右侧指针-迭代O(1)", connect3(root))
}

// connect 层序遍历，每次只添加每一层的最后一个元素
func connect(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	levelNodeQueue := []*TreeNode{root}
	for len(levelNodeQueue) > 0 {
		length := len(levelNodeQueue)

		var pre *TreeNode
		for j := 0; j < length; j++ {
			node := levelNodeQueue[0]
			levelNodeQueue = levelNodeQueue[1:]

			if pre != nil {
				pre.Next = node
			}

			pre = node

			// 入队列，按照从左到右的顺序
			if node.Left != nil {
				levelNodeQueue = append(levelNodeQueue, node.Left)
			}

			if node.Right != nil {
				levelNodeQueue = append(levelNodeQueue, node.Right)
			}
		}

		pre = nil
	}

	return root
}

// connect2 递归的过程没法判断当前是不是最后一个结点，只能层序遍历完再把最后一个元素取出来
func connect2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var modify func(root *TreeNode) *TreeNode

	modify = func(root *TreeNode) *TreeNode {
		if root.Left == nil && root.Right == nil {
			return root
		}
		root.Left.Next = root.Right
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
		modify(root.Left)
		modify(root.Right)
		return root
	}

	return modify(root)
}

// 时间复杂度O(N),空间复杂度O(1)
func connect3(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var nextLevel *TreeNode
	cur := root

	for cur.Left != nil {
		nextLevel = cur.Left
		for cur != nil {
			// 处理左节点指向右结点
			cur.Left.Next = cur.Right

			// 处理右节点指向跨级结点
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}

			// 横着平移，直到这一层遍历完毕
			cur = cur.Next
		}

		// 从每一层的左边开始循环
		cur = nextLevel
	}

	return root
}
