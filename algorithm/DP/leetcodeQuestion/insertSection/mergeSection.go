package main

import (
	"log"
	"sort"
)

/**
题目：https://leetcode-cn.com/problems/insert-interval/

插入区间
给你一个 无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。


示例1：

输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]
示例 2：

输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10]重叠。
示例 3：

输入：intervals = [], newInterval = [5,7]
输出：[[5,7]]
示例 4：

输入：intervals = [[1,5]], newInterval = [2,3]
输出：[[1,5]]
示例 5：

输入：intervals = [[1,5]], newInterval = [2,7]
输出：[[1,7]]

提示：

0 <= intervals.length <= 104
intervals[i].length == 2
0 <=intervals[i][0] <=intervals[i][1] <= 105
intervals 根据 intervals[i][0] 按 升序 排列
newInterval.length == 2
0 <=newInterval[0] <=newInterval[1] <= 105


*/

func main() {
	intervals := [][]int{{1, 3}, {8, 10}, {15, 18}}
	newInterval := []int{2, 6}
	log.Println("合并区间-排序合并:", insert(intervals, newInterval))
	log.Println("合并区间-直接模拟:", insert2(intervals, newInterval))
}

// insert 先排序后插入 O(nlogn)
func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	intervals = append(intervals, newInterval)
	// 因为区间都是有序的，所以先把区间按照开头元素排好序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		// intervals[i]至少是2个元素，所以可以不用判断直接写
		if intervals[i][0] <= end {
			if intervals[i][1] > end {
				// 说明2和区间可以合并，第一个区间的start,第二个区间的end
				end = intervals[i][1]
			}
		} else {
			// 不能合并的话，直接加到结果集，向前移动
			res = append(res, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		}
	}

	// 最后一个区间也得加上
	return append(res, []int{start, end})
}

// insert2 直接模拟 O(n)
func insert2(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	l := len(intervals)
	i := 0
	for i < l && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	for i < l && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	res = append(res, newInterval)
	for i < l {
		res = append(res, intervals[i])
		i++
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
