package removeduplicate

import "fmt"

/*
	给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

	不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

	大前提是已经排序过了：
*/

func DoRemove() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//l := removeDuplicates(nums)
	l := removeDuplicates1(nums)
	fmt.Println("l,nums:", l, nums)
}

func removeDuplicates(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}

func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 1
	preNum := nums[0]
	for _, v := range nums {
		if v != preNum {
			preNum = v
			nums[count] = v
			count += 1
		}
	}
	return count
}

func removeDuplicates2(nums []int) int {
	nui := 0
	for i := 1; i < len(nums); i++ {
		if nums[nui] != nums[i] {
			nui++
			nums[nui] = nums[i]
		}
	}
	return nui + 1
}

/**
i超出范围会报越界错误 runtime error: slice bounds out of range
1.nums[:i]		---// 冒号在左边，从左往右保留i个
2.nums[i:]		---// 冒号在右边，从左往右删除i个元素
3.nums[:], nums[0:]	等价于复制整个数组元素
*/
func TTT() {
	nums := []int{0, 1, 2, 3, 4, 5}
	// i表示要删除索引的位置
	i := 2
	fmt.Println("nums[:i]:", nums[:i])   //删除索引的前半部分
	fmt.Println("nums[:i]:", nums[i+1:]) //删除索引的后半部分
	fmt.Println("nums[i:]:", nums[:], nums[0:])
}
