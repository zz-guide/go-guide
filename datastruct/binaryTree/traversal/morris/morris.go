package morris

import (
	. "go-guide/datastruct/binaryTree/treeNode"
)

/**
Morris序转化前中后序遍历：
	前序遍历：1.没有Left，直接打印。2.第一次遍历的结点，直接打印。3.打印的都是cur的值
	中序遍历：1.对于能第二次遍历到的结点，第二次打印。2.第一次遍历的结点，直接打印。3.打印的都是cur的值
	后序遍历：1.能两次遍历到的结点，并且是第二次遍历的时候打印。2.逆序打印这些结点的右边界。3.最后再加上整颗树的右边界即可。
	子树的右边界可以分界整颗树，后序思想。
	是不是搜索二叉树？1.root是nil就是
	4.逆序打印需要反转单链表，线索二叉树
*/

// TraversalMorris Morris遍历
// 步骤：1.从root开始遍历
// 2.寻找当前结点cur的左子树的最右子树
// 3.若当前结点cur没有左子树，则向右移动
// 4.如果找到最右子树，判断是不是nil，如果是nil则证明是首次遍历，mostRight.Right = cur,然后向左移动重复此过程
// 原理：利用二叉树某些Node的空闲指针可以回到上层结点
// 时间复杂度：O(N)，空间复杂度：O(1)，就2个变量
func TraversalMorris(root *TreeNode) ([]int, []int, []int, []int, bool) {
	var isBST bool = false // 是不是搜索二叉树

	if root == nil {
		return nil, nil, nil, nil, true
	}

	var morrisRes []int    // 存放cur每次遍历ur结点的值
	var preorderRes []int  // 存放前序遍历
	var inorderRes []int   // 存放中序遍历
	var postorderRes []int // 存放后序遍历

	var mostRight *TreeNode // 存放左子树的最右子树结点
	cur := root
	var curVal *int

	// 等同于while循环，cur为nil退出循环，也就遍历完了
	for cur != nil {

		morrisRes = append(morrisRes, cur.Val)

		// 首先把cur的左子树赋值给mostRight
		mostRight = cur.Left
		if mostRight != nil { // 存在左子树的逻辑
			// 不停的寻找最右子树，mostRight.Right != cur用来判断是首次遍历
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}

			// 首次遍历到这个最右结点，更改指向，并向左移动
			if mostRight.Right == nil {
				//收集前序遍历的值
				preorderRes = append(preorderRes, cur.Val)
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else { // 再次遍历到最右结点，恢复nil指向，向右移动
				if curVal != nil && *curVal > cur.Val {
					isBST = false
				}
				// 中序遍历
				inorderRes = append(inorderRes, cur.Val)

				mostRight.Right = nil

				// 后序遍历
				postorderRes = append(postorderRes, getRightEdge(cur.Left)...)

				// 这2行其实可以不要，因为 最外层else也是右移
				cur = cur.Right
				continue
			}

		} else {
			// 收集前序遍历的值
			preorderRes = append(preorderRes, cur.Val)
			if curVal != nil && *curVal > cur.Val {
				isBST = false
			}

			// 收集中序遍历的值
			inorderRes = append(inorderRes, cur.Val)

			cur = cur.Right // 不存在左子树就右移,注意：此处可能为空
		}

		if cur != nil {
			curVal = &cur.Val
		}
	}

	postorderRes = append(postorderRes, getRightEdge(root)...)
	//postorderRes = reverseSlice(postorderRes)
	return morrisRes, preorderRes, inorderRes, postorderRes, isBST
}

// getRightEdge 获取右边界
func getRightEdge(root *TreeNode) []int {
	var res []int
	tail := reverseRightSingleLinked(root)
	cur := tail
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Right
	}

	// 恢复指向
	reverseRightSingleLinked(tail)
	return res
}

// 逆序右边界
func reverseRightSingleLinked(root *TreeNode) *TreeNode {
	var pre *TreeNode
	var next *TreeNode
	for root != nil {
		next = root.Right // 下一次反转的结点
		root.Right = pre
		pre = root
		root = next
	}
	return pre
}
