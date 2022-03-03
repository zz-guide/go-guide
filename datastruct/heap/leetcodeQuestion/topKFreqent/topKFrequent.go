package main

import (
	"container/heap"
	"log"
	"math/rand"
	"sort"
	"time"
)

/**
题目：https://leetcode-cn.com/problems/top-k-frequent-elements/submissions/

前 K 个高频元素

给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

提示：

1 <= nums.length <= 105
k 的取值范围是 [1, 数组中不相同的元素的个数]
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的


进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n是数组大小。

注意：
1.结果只要前k个，与顺序无关，可以不用排序
2.堆排序不需要维护N个元素，只维护K个即可
3.快速排序，只需要排一个分支即可


*/
func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 6, 7, 8}
	k := 4
	log.Println("前K个高频元素-堆:", topKFrequent(nums, k))
	log.Println("前K个高频元素-快速排序:", topKFrequent1(nums, k))
	log.Println("前K个高频元素-哈希:", topKFrequent2(nums, k))
	log.Println("前K个高频元素-改进快排:", topKFrequent3(nums, k))
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	length := h.Len()
	x := (*h)[length-1]
	*h = (*h)[0 : length-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	freqMap := map[int]int{}
	for _, num := range nums {
		freqMap[num]++
	}

	// 初始化一个堆
	h := &IHeap{}
	heap.Init(h)
	// 不满k的时候一直加，满了以后比较当前元素和堆顶哪个大，留哪个，堆里只保留k个元素
	for v, freq := range freqMap {
		heap.Push(h, [2]int{v, freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 出堆
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[k-i-1] = (heap.Pop(h).([2]int))[0]
	}

	return res
}

// topKFrequent1
func topKFrequent1(nums []int, k int) []int {
	freqMap := make(map[int]int) // 存放值->出现频率
	valueArray := make([]int, 0) // 存放第一次出现的值
	for _, v := range nums {
		i, ok := freqMap[v]
		if ok {
			freqMap[v] = i + 1
		} else {
			freqMap[v] = 1
			valueArray = append(valueArray, v)
		}
	}

	// 提供freq作为排序依据
	sort.Slice(valueArray, func(i, j int) bool {
		return freqMap[valueArray[i]] > freqMap[valueArray[j]]
	})

	// 返回前k个
	return valueArray[:k]
}

// topKFrequent2 纯hash， O(n^2)
func topKFrequent2(nums []int, k int) []int {
	// 值->频率
	m1 := map[int]int{}
	for _, v := range nums {
		m1[v]++

	}

	// 频率->值 数组
	m2 := map[int][]int{}
	max := 0
	for k, v := range m1 {
		m2[v] = append(m2[v], k)
		if v > max {
			max = v
		}
	}

	var res []int
	for i := max; i >= 0 && k > 0; i-- {
		if v, ok := m2[i]; ok {
			if len(v) >= k {
				res = append(res, v[0:k]...)
				k = 0
			} else {
				res = append(res, v...)
				k -= len(v)
			}
		}
	}

	return res
}

// topKFrequent3 快速排序改进版，只需要排一边即可(没理解，先放着)
func topKFrequent3(nums []int, k int) []int {
	freqMap := map[int]int{}
	for _, num := range nums {
		freqMap[num]++
	}

	var values [][]int
	for v, freq := range freqMap {
		values = append(values, []int{v, freq})
	}

	ret := make([]int, k)
	quickSort(values, 0, len(values)-1, ret, 0, k)
	return ret
}

func quickSort(values [][]int, start, end int, ret []int, retIndex, k int) {
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int()%(end-start+1) + start
	values[picked], values[start] = values[start], values[picked]

	pivot := values[start][1]
	index := start

	for i := start + 1; i <= end; i++ {
		if values[i][1] >= pivot {
			values[index+1], values[i] = values[i], values[index+1]
			index++
		}
	}
	values[start], values[index] = values[index], values[start]
	if k <= index-start {
		quickSort(values, start, index-1, ret, retIndex, k)
	} else {
		for i := start; i <= index; i++ {
			ret[retIndex] = values[i][0]
			retIndex++
		}

		if k > index-start+1 {
			quickSort(values, index+1, end, ret, retIndex, k-(index-start+1))
		}
	}
}
