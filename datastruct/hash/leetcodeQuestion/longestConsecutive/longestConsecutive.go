package main

import (
	"log"
)

/**
题目：https://leetcode-cn.com/problems/longest-consecutive-sequence/

最长连续序列

给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为O(n) 的算法解决此问题。

示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9

提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109

*/
func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	log.Println("最长连续序列-哈希:", longestConsecutive(nums))
	log.Println("最长连续序列-并查集:", longestConsecutive2(nums))
}

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	res := 0
	for num := range numSet {
		// 先判断num-1有没有，没有则从当前数字开始增加，一直判断
		// +1，-1 都行，取决于是从大到小寻扎还是从小到大寻找
		if !numSet[num-1] {
			currentNum := num
			cnt := 1
			for numSet[currentNum+1] {
				currentNum++
				cnt++
			}

			// 交换结果
			if res < cnt {
				res = cnt
			}
		}
	}

	return res
}

// longestConsecutive2 并查集
func longestConsecutive2(nums []int) int {
	uf := NewUF(nums)
	res := 0
	for i := 0; i < len(nums); i++ { //遍历数组
		n, tmp := nums[i], nums[i] //tmp记录原始值
		for {
			if _, ok := uf.Find(n - 1); !ok { //如果不存在比当前值-1的值，则肯定不连通
				break
			}

			if !uf.Connected(n-1, n) { //如果存在n-1并且还没有连通，则连一下
				uf.Union(n-1, n)
				n = n - 1
			} else { //如果n-1与n是连通的，则求一下父节点值
				n, _ = uf.Find(n)
				break
			}
		}
		res = max(res, tmp-n+1) //记录一下，连同节点的最大距离即为最长序列
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type UF struct {
	parent map[int]int //每个节点的父节点
}

func NewUF(nums []int) *UF {
	uf := UF{}
	uf.parent = make(map[int]int)
	for _, num := range nums {
		uf.parent[num] = num
	}
	return &uf
}

func (uf *UF) Connected(a, b int) bool {
	pA, _ := uf.Find(a)
	pB, _ := uf.Find(b)
	return pA == pB
}

func (uf *UF) Union(a, b int) {
	pA, _ := uf.Find(a)
	pB, _ := uf.Find(b)
	if pA == pB {
		return
	}
	if pA > pB { //将值小的节点作为父节点
		uf.parent[pA] = pB
	} else {
		uf.parent[pB] = pA
	}
}

func (uf *UF) Find(a int) (int, bool) {
	if _, ok := uf.parent[a]; !ok {
		return -1, false
	}

	for a != uf.parent[a] {
		uf.parent[a] = uf.parent[uf.parent[a]] //压缩树的高度
		a = uf.parent[a]
	}

	return a, true
}
