package main

import (
	. "go-guide/datastruct/list/node"
	"log"
)

/**
题目：https://leetcode.cn/problems/rotate-list/
旋转链表
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
示例 1：

输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]

提示：

链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 109


*/
func main() {
	head := MakeListNode([]int{1, 2, 3, 5, 7})
	k := 2
	log.Println("旋转链表", rotateRight(head, k))
}

func rotateRight(head *ListNode, k int) *ListNode {
	return nil
}
