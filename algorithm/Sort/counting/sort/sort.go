package main

import (
	"log"
	"math/rand"
	"time"
)

/**
计数排序是一个非基于比较的排序算法，该算法于1954年由 Harold H. Seward 提出。
它的优势在于在对一定范围内的整数排序时，它的复杂度为Ο(n+k)（其中k是整数的范围），
快于任何比较排序算法。 [1]  当然这是一种牺牲空间换取时间的做法，
而且当O(k)>O(n*log(n))的时候其效率反而不如基于比较的排序
（基于比较的排序的时间复杂度在理论上的下限是O(n*log(n)), 如归并排序，堆排序）

复杂度
最佳情况：T(n) = O(n+k)
最坏情况：T(n) = O(n+k)
平均情况：T(n) = O(n+k)
空间复杂度：O(k)
稳定性：稳定
排序方式：Out-place

*/

func init() {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
}

func main() {
	slice := make([]int, 0)
	// 生成100个1000以内的随机数
	for i := 1; i <= 10; i++ {
		slice = append(slice, rand.Intn(1000))
	}

	CountingSortAsc(slice)
	log.Println("计数排序:", slice)
}

// CountingSortAsc 计数排序--升序
// 找出数组的最大值和最小值
// 创建新数组[max+1]， 原数组值等于新数组下标时， 新数组值加一
// 最后把新数组下标按值的个数输出即排序完成
func CountingSortAsc(nums []int) {
	min, max := CountMaxMin(nums)
	res := make([]int, max+1) // 从0开始，所以需要多分配一个
	for i := 0; i < len(nums); i++ {
		res[nums[i]]++
	}

	var index int // 标记新的结果覆盖到哪了
	for i := min; i < len(res); i++ {
		// 可能一个位置有多个相等的值
		for j := res[i]; j > 0; j-- {
			// 从0开始覆盖
			nums[index] = i
			index++
		}
	}
}

func CountMaxMin(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}

		if max < nums[i] {
			max = nums[i]
		}
	}

	return min, max
}
