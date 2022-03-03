package main

import (
	"container/heap"
	"log"
	"sort"
)

/**
题目：https://leetcode-cn.com/problems/longest-happy-string/
最长快乐字符串

如果字符串中不含有任何 'aaa'，'bbb' 或 'ccc' 这样的字符串作为子串，那么该字符串就是一个「快乐字符串」。

给你三个整数 a，b ，c，请你返回 任意一个 满足下列全部条件的字符串 s：

s 是一个尽可能长的快乐字符串。
s 中 最多 有a 个字母 'a'、b个字母 'b'、c 个字母 'c' 。
s 中只含有 'a'、'b' 、'c' 三种字母。
如果不存在这样的字符串 s ，请返回一个空字符串 ""。

示例 1：

输入：a = 1, b = 1, c = 7
输出："ccaccbcc"
解释："ccbccacc" 也是一种正确答案。
示例 2：

输入：a = 2, b = 2, c = 1
输出："aabbc"
示例 3：

输入：a = 7, b = 1, c = 0
输出："aabaa"
解释：这是该测试用例的唯一正确答案。

提示：

0 <= a, b, c <= 100
a + b + c > 0

*/
func main() {
	a := 2
	b := 2
	c := 1
	log.Println("最长快乐字符串-贪心算法:", longestDiverseString(a, b, c))
	log.Println("最长快乐字符串-堆排序:", longestDiverseString1(a, b, c))
}

type pair struct {
	b byte
	c int
}

// longestDiverseString 贪心算法，时间复杂度：O((a+b+c)×ClogC)，空间复杂度：O(C)
func longestDiverseString(a int, b int, c int) string {
	var res string
	pairs := []pair{{'a', a}, {'b', b}, {'c', c}}
Loop:
	for {
		// 排序，把剩余字符数最多的放在第一位
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].c > pairs[j].c
		})

		for i := range pairs {
			// 次数为0说明不能取了，直接返回
			if pairs[i].c == 0 {
				break Loop
			}

			// 1.长度小于2，不管什么字符，肯定不是快乐数
			// 2.res倒数2个字符不等，如果相等，有可能是快乐数
			// 3.当前字符不等于res最后一个字符
			if len(res) < 2 || res[len(res)-1] != res[len(res)-2] || res[len(res)-1] != pairs[i].b {
				res += string(pairs[i].b)
				pairs[i].c--
				break
			}
		}
	}

	return res
}

// longestDiverseString1 堆排序，时间复杂度：令答案最大长度为 n=a+b+c，优先队列中最多有 C = 3个元素，复杂度为 O(n∗k∗logC)，其中 k 为构造答案字符串中每个字符所需要的平均「出队 + 入队」次数，k 为一个范围在 [2,4] 的数字
func longestDiverseString1(a int, b int, c int) string {
	myHp := &hp{}
	if a > 0 {
		heap.Push(myHp, node{-a, 'a'})
	}
	if b > 0 {
		heap.Push(myHp, node{-b, 'b'})
	}
	if c > 0 {
		heap.Push(myHp, node{-c, 'c'})
	}
	var tmpq []node
	var ans []byte
	for myHp.Len() > 0 {
		top := heap.Pop(myHp).(node)
		if top.num < (-2) && (len(tmpq) == 0 || top.num < tmpq[0].num) {
			ans = append(ans, []byte{top.cha, top.cha}...)
			top.num += 2
		} else {
			ans = append(ans, top.cha)
			top.num++
		}
		tmpq = append(tmpq, node{top.num, top.cha})
		for len(tmpq) >= 2 {
			cur := tmpq[0]
			tmpq = tmpq[1:]
			if cur.num < 0 {
				heap.Push(myHp, cur)
			}
		}
	}
	if len(tmpq) > 0 {
		cur := tmpq[0]
		if cur.num < 0 && cur.cha != ans[len(ans)-1] {
			ans = append(ans, tmpq[0].cha)
		}
	}
	return string(ans)
}

type node struct {
	num int
	cha byte
}
type hp []node

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].num < h[j].num }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Pop() interface{} {
	a := *h
	n := len(a)
	v := a[n-1]
	*h = a[:n-1]
	return v
}
func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(node))
}
