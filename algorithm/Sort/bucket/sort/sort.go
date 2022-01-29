package main

import (
	"log"
)

/**
桶排序 (Bucket sort)或所谓的箱排序，是一个排序算法，工作的原理是将数组分到有限数量的桶子里。每个桶子再个别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排序）。桶排序是鸽巢排序的一种归纳结果。当要被排序的数组内的数值是均匀分配的时候，桶排序使用线性时间（Θ（n））。但桶排序并不是 比较排序，他不受到 O(n log n) 下限的影响。
桶排序的平均时间复杂度为线性的O(N+C)，其中C=N*(logN-logM)。如果相对于同样的N，桶数量M越大，其效率越高，最好的时间复杂度达到O(N)。当然桶排序的空间复杂度为O(N+M)，如果输入数据非常庞大，而桶的数量也非常多，则空间代价无疑是昂贵的。此外，桶排序是稳定的。
*/
func main() {
	nums := []int{10, 1, 18, 30, 23, 12, 7, 5, 18, 17}
	log.Println("桶排序:", bucketSort(nums))
}

/*
桶内排序
*/
func sortInBucket(bucket []int) { //此处实现插入排序方式，其实可以用任意其他排序方式
	length := len(bucket)
	if length == 1 {
		return
	}

	for i := 1; i < length; i++ {
		backup := bucket[i]
		j := i - 1
		//将选出的被排数比较后插入左边有序区
		for j >= 0 && backup < bucket[j] { //注意j >= 0必须在前边，否则会数组越界
			bucket[j+1] = bucket[j] //移动有序数组
			j--                     //反向移动下标
		}
		bucket[j+1] = backup //插队插入移动后的空位
	}
}

/*
获取数组最大值
*/
func getMaxInArr(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

/*
桶排序
*/
func bucketSort(nums []int) []int {
	//桶数
	num := len(nums)
	//k（数组最大值）
	max := getMaxInArr(nums)
	//二维切片
	buckets := make([][]int, num)

	//分配入桶
	index := 0
	for i := 0; i < num; i++ {
		index = nums[i] * (num - 1) / max //分配桶index = value * (n-1) /k

		buckets[index] = append(buckets[index], nums[i])
	}
	//桶内排序
	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			sortInBucket(buckets[i])

			copy(nums[tmpPos:], buckets[i])

			tmpPos += bucketLen
		}
	}

	return nums
}
