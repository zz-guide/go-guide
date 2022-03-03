package main

import (
	. "go-guide/datastruct/list/node"
)

/**
题目：https://leetcode-cn.com/problems/sort-list/

排序链表
给你链表的头结点head，请将其按 升序 排列并返回 排序后的链表 。

示例 1：

输入：head = [4,2,1,3]
输出：[1,2,3,4]
示例 2：


输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]
示例 3：

输入：head = []
输出：[]


提示：

链表中节点的数目在范围[0, 5 * 104]内
-105<= Node.val <= 105


进阶：你可以在O(nlogn) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？


注意：1.O(nlogn)时间复杂度，只能是迭代版本的排序了，比如快速排序，归并排序

*/

func main() {
	list1 := MakeListNode([]int{1, 3, 2, 7, 5})
	//PrintListNode(sortList(list1))  // 迭代归并
	//PrintListNode(sortList1(list1)) // 自顶向下归并
	PrintListNode(sortList3(list1)) // 自顶向下归并
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// 计算整个链表的长度
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummyHead := &ListNode{Next: head}

	step := 1 // 步长间隔为1
	for step < length {
		for i := 0; i+step < length; i += step * 2 {

		}
	}

	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = merge(head1, head2)
			for prev.Next != nil {
				prev = prev.Next
			}

			cur = next
		}
	}

	return dummyHead.Next
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	// 快慢指针移动把链表分为2部分，slow指向的就是中点
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func sortList1(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sortList3(head *ListNode) *ListNode {
	// 因为可能涉及到对头节点进行操作，所以设置一个dummy, 来减少边界判断
	dummy := &ListNode{}
	dummy.Next = head

	// 先求链表长度
	p := head
	n := 0
	for p != nil {
		p = p.Next
		n++
	}

	// 用递归进行归并排序的时候，递归最底层其实是将节点两两分组（每组两个节点），进行排序，然后逐段向上归并
	// 所以在手动迭代的时候也我们先把所有节点两两分组，然后再四四分组.. 从两两分组开始排序
	for i := 1; i < n; i = i * 2 { // i表示每一小段的节点个数，i = 1的时候表示将节点两两分组
		cur := dummy                       // 每一个i循环，表示归并中的每一层操作, 所以cur每次都要从头开始
		for j := 0; j+i < n; j = j + 2*i { // j表示一共有多少个小组要进行排序操作，每一组的个数是 2 * i个
			l, r := 0, 0                      // 用l, r分别记录需要排序的左右两小段各并入了几个点
			left, right := cur.Next, cur.Next // left和right分别指向要排序的两段的第一个位置
			for k := 0; k < i; k++ {          // right移动到第二段的第一个位置
				right = right.Next
			}
			// 开始归并排序, 这个过程不要想的太复杂，那个节点的值比较小，那就说明它要接到cur的后面, 可以用两个节点模拟一下
			for l < i && r < i && right != nil {
				// right这里要判断一下不为空，因为如果链表总长为奇数个，最后一段会不足i个，处理i次之后就会空了
				if left.Val <= right.Val {
					cur.Next = left
					cur = cur.Next
					left = left.Next
					l += 1
				} else {
					cur.Next = right
					cur = cur.Next
					right = right.Next
					r += 1
				}
			}
			// 退出循环的时候一定有一段已经处理完了，那就处理另外一段
			for l < i {
				cur.Next = left
				cur = cur.Next
				left = left.Next
				l += 1
			}
			for r < i && right != nil { // 同理判断right不为空
				cur.Next = right
				cur = cur.Next
				right = right.Next
				r += 1
			}

			// 这样一小段的归并就处理完了，也就是处理了2*i个节点
			// 每一段处理完后right一定是停在下一段的开始位置，这个时候要让cur.Next指向r
			// 否则最后一个如果处理的是左半边的点会出现死循环
			cur.Next = right
		}
	} // 一个j循环结束就表示这一层的节点都按照2 * i的分组规矩处理完成了，以2*i为单位的每一小段都是有序了
	// 接下来将i翻倍，继续同样的处理，直到处理完i = n / 2
	return dummy.Next
}
