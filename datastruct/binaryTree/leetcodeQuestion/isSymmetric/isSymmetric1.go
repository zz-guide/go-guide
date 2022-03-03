package main

import (
	. "go-guide/datastruct/binaryTree/treeNode"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/symmetric-tree/
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树[1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3


但是下面这个[1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

注意：对称其实就是指根节点的2颗子树是不是一方反转之后与另外一方相等

*/
func main() {
	root := NewSymmetricTree()

	log.Println("是不是对称二叉树-dfs迭代法:", isSymmetric1(root))
	log.Println("是不是对称二叉树-dfs递归:", isSymmetric2(root))
	log.Println("是不是对称二叉树-bfs迭代:", isSymmetric3(root))
}

// isSymmetric1 dfs迭代法,复制一份root,同时比较2个root,但是方向不一样
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l, r := root, root
	// 队列每次存放下一次需要比较的元素，按照左右，右左的顺序放进去进行比较，最开始就是根元素自身
	queue := []*TreeNode{l, r}

	for len(queue) > 0 {
		// 出队列
		l, r = queue[0], queue[1]
		queue = queue[2:]

		// 如果左节点和右节点都是nil,直接进行下一次循环
		if l == nil && r == nil {
			continue
		}

		// 其中有一个nil代表不符合条件，肯定不对称
		if l == nil || r == nil {
			return false
		}

		// 两个节点值一样代表是对称
		if l.Val != r.Val {
			return false
		}

		// 按照对称的比较顺序入队列
		queue = append(queue, l.Left, r.Right, l.Right, r.Left)
	}

	return true
}

// isSymmetric2 递归法 时间复杂度O(n)，空间复杂度为O(n)
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var check func(p, q *TreeNode) bool
	check = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}

		return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
	}

	return check(root.Left, root.Right)
}

// isSymmetric3 层序遍历bfs迭代，需要把头结点过滤掉，其次，每一层的结点即使有nil的也需要填充值
func isSymmetric3(root *TreeNode) bool {
	if root == nil {
		return true
	}

	levelNodeQueue := []*TreeNode{root.Left, root.Right}

	for len(levelNodeQueue) > 0 {
		// 进入新的一层需要留出一个位置放结果
		var ret []*TreeNode
		length := len(levelNodeQueue)
		for j := 0; j < length; j++ {
			// 出栈，每次取第一个元素
			node := levelNodeQueue[0]
			levelNodeQueue = levelNodeQueue[1:]

			if node != nil {
				ret = append(ret, node)
				// 入队列，按照从左到右的顺序, nil也必须入，判断对称需要的是左右结点全部的信息
				levelNodeQueue = append(levelNodeQueue, node.Left)
				levelNodeQueue = append(levelNodeQueue, node.Right)
			} else {
				ret = append(ret, nil)
			}
		}

		// 如果是奇数，肯定不是对称的
		if length%2 != 0 {
			return false
		}

		for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
			if ret[i] == nil && ret[j] == nil {
				continue
			}

			if ret[i] == nil || ret[j] == nil {
				return false
			}

			if ret[i].Val != ret[j].Val {
				return false
			}
		}
	}

	return true
}
