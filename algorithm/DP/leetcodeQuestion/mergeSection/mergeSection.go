package main

import (
	"log"
	"sort"
)

/**
题目：https://leetcode-cn.com/problems/merge-intervals/

合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。


示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

提示：

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104

*/

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	log.Println("合并区间:", merge(intervals))
}

func merge(intervals [][]int) [][]int {
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
