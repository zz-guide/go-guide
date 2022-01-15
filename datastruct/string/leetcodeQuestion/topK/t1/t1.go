package main

import (
	"container/heap"
	"sort"
)

/**
一非空的单词列表，返回前k个出现次数最多的单词。

返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。

示例 1：

输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
输出: ["i", "love"]
解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
    注意，按字母顺序 "i" 在 "love" 之前。

示例 2：

输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
输出: ["the", "is", "sunny", "day"]
解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
    出现次数依次为 4, 3, 2 和 1 次。

注意：

假定 k 总为有效值， 1 ≤ k ≤ 集合元素数。
输入的单词均由小写字母组成。

扩展练习：

尝试以O(n log k) 时间复杂度和O(n) 空间复杂度解决。

*/
func main() {

}

/**
哈希表 + 排序
复杂度分析

时间复杂度：O(l \times n + l \times m \log m)O(l×n+l×mlogm)，其中 nn 表示给定字符串序列的长度，ll 表示字符串的平均长度，mm 表示实际字符串种类数。我们需要 l \times nl×n 的时间将字符串插入到哈希表中，以及 l \times m \log ml×mlogm 的时间完成字符串比较（最坏情况下所有字符串出现频率都相同，我们需要将它们两两比较）。

空间复杂度：O(l \times m)O(l×m)，其中 ll 表示字符串的平均长度，mm 表示实际字符串种类数。哈希表和生成的排序数组空间占用均为 O(l \times m)O(l×m)。
*/
func topKFrequent(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	uniqueWords := make([]string, 0, len(cnt))
	for w := range cnt {
		uniqueWords = append(uniqueWords, w)
	}
	sort.Slice(uniqueWords, func(i, j int) bool {
		s, t := uniqueWords[i], uniqueWords[j]
		return cnt[s] > cnt[t] || cnt[s] == cnt[t] && s < t
	})
	return uniqueWords[:k]
}

type pair struct {
	w string
	c int
}
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { a, b := h[i], h[j]; return a.c < b.c || a.c == b.c && a.w > b.w }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

/**
优先队列
复杂度分析

时间复杂度：O(l \times n + m \times l \log k)O(l×n+m×llogk)，其中 nn 表示给定字符串序列的长度，mm 表示实际字符串种类数，ll 表示字符串的平均长度。我们需要 l \times nl×n 的时间将字符串插入到哈希表中，以及每次插入元素到优先队列中都需要 l \log kllogk 的时间，共需要插入 mm 次。

空间复杂度：O(l \times (m + k))O(l×(m+k))，其中 ll 表示字符串的平均长度，mm 表示实际字符串种类数。哈希表空间占用为 O(l \times m)O(l×m)，优先队列空间占用为 O(l \times k)O(l×k)。
*/
func topKFrequent1(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	h := &hp{}
	for w, c := range cnt {
		heap.Push(h, pair{w, c})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(pair).w
	}
	return ans
}
