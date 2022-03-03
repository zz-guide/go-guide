package main

import (
	"go-guide/datastruct/binaryTree/traversal/levelorder"
	. "go-guide/datastruct/binaryTree/treeNode"
	. "go-guide/datastruct/list/node"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/solution/you-xu-lian-biao-zhuan-huan-er-cha-sou-suo-shu-1-3/

有序链表转换二叉搜索树

给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

*/

func main() {
	list1 := MakeListNode([]int{-10, -3, 0, 5, 9})
	log.Println("有序链表转换二叉搜索树-二分查找中序遍历:", levelorder.TraversalRecursive(sortedListToBST(list1)))
	log.Println("有序链表转换二叉搜索树-快慢指针:", levelorder.TraversalRecursive(sortedListToBST2(list1)))
}

// sortedListToBST O(n) O(logn)
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	// 先计算链表的长度
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}

	var buildTree func(left, right int) *TreeNode
	buildTree = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		// 和数组不一样，链表没办法直接定位到根结点，需要先把根节点左半部分全部处理完才能处理根结点自身
		mid := (left + right + 1) / 2
		// 先处理根结点左半部分
		root := &TreeNode{}
		root.Left = buildTree(left, mid-1)
		// 处理自身，不能提前赋值
		root.Val = head.Val
		head = head.Next
		// 处理右半部分
		root.Right = buildTree(mid+1, right)
		return root
	}

	return buildTree(0, length-1)
}

// 快慢指针 sortedListToBST2
func sortedListToBST2(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

func getMedian(left, right *ListNode) *ListNode {
	fast, slow := left, left
	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getMedian(left, right)
	root := &TreeNode{Val: mid.Val}
	root.Left = buildTree(left, mid)
	root.Right = buildTree(mid.Next, right)
	return root
}
