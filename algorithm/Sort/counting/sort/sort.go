package main

import (
	"log"
	"math/rand"
	"time"
)

/**
计数排序是一个非基于比较的排序算法，该算法于1954年由 Harold H. Seward 提出。它的优势在于在对一定范围内的整数排序时，它的复杂度为Ο(n+k)（其中k是整数的范围），快于任何比较排序算法。 [1]  当然这是一种牺牲空间换取时间的做法，而且当O(k)>O(n*log(n))的时候其效率反而不如基于比较的排序（基于比较的排序的时间复杂度在理论上的下限是O(n*log(n)), 如归并排序，堆排序）

*/
func main() {
	slice := make([]int, 0)
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	// 生成100个1000以内的随机数
	for i := 1; i <= 10; i++ {
		slice = append(slice, rand.Intn(1000))
	}

	ContingSortAsc(slice)
	log.Println("计数排序升序:", slice)
	ContingSortDesc(slice)
	log.Println("计数排序降序:", slice)
}

// 计数排序--升序
func ContingSortAsc(slice []int) {
	// 创建map统计0-999每个数出现的次数
	m := make(map[int]int)
	// 遍历待排序的数据，统计结果
	for _, v := range slice {
		m[v]++
	}
	// 借助map，统计排序的数据重新赋值为原序列
	slice = slice[0:0] // 将原序列清空
	for i := 0; i < 1000; i++ {
		// for i := 0; i < 1000; i++ {
		// 数据出现的次数：m[i]的值
		for j := 0; j < m[i]; j++ {
			slice = append(slice, i) // 重新赋值
		}
	}
}

// 计数排序--降序
func ContingSortDesc(slice []int) {
	// 创建map统计0-999每个数出现的次数
	m := make(map[int]int)
	// 遍历待排序的数据，统计结果
	for _, v := range slice {
		m[v]++
	}
	// 借助map，统计排序的数据重新赋值为原序列
	slice = slice[0:0] // 将原序列清空
	for i := 999; i >= 0; i-- {
		// for i := 0; i < 1000; i++ {
		// 数据出现的次数：m[i]的值
		for j := 0; j < m[i]; j++ {
			slice = append(slice, i) // 重新赋值
		}
	}
}
