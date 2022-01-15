package node

import "fmt"

// ListNode 单链表结构
type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{
		Val: nums[0],
	}

	temp := head

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val: nums[i]}
		temp = temp.Next
	}

	return head
}

func MakeCircleListNode(nums []int, firstIndex, tailIndex int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{
		Val: nums[0],
	}

	temp := head
	var firstNode *ListNode
	for i := 1; i < len(nums); i++ {
		tempNode := &ListNode{Val: nums[i]}
		temp.Next = tempNode
		if i == firstIndex {
			firstNode = tempNode
		}

		if i == tailIndex {
			tempNode.Next = firstNode
			break
		}

		temp = temp.Next
	}

	return head
}

func MakeListNodeByNode(nums []int, node *ListNode) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{
		Val: nums[0],
	}

	temp := head

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val: nums[i]}
		temp = temp.Next
	}

	temp.Next = node
	return head
}

func PrintListNode(node *ListNode) {
	temp := node
	res := make([]int, 0)
	for temp != nil {
		res = append(res, temp.Val)
		temp = temp.Next
	}

	fmt.Println("result:", res)
}

func PrintSingleListNode(node *ListNode) {
	if node != nil {
		fmt.Println("node:", node.Val)
		return
	}

	fmt.Println("node:", "")
}
