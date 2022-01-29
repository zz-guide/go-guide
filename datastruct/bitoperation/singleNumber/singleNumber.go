package main

import "log"

/**
题目：https://leetcode-cn.com/problems/single-number/solution/
只出现一次的数字

给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例2:

输入: [4,1,2,1,2]
输出: 4

注意：1.O(1)时间复杂度

异或运算的规则：
1.任何数和 0 做异或运算，结果仍然是原来的数，
2.任何数和其自身做异或运算，结果是 0
3.异或运算满足交换律和结合律：a⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=b

异或：1.相同为0，不同为1，这些法则与加法是相同的，只是不带进位，所以异或常被认作不进位加法。

*/
func main() {
	nums := []int{4, 1, 2, 1, 2}
	log.Println("只出现一次的数字:", singleNumber(nums))
}

// singleNumber 时间复杂度O(1),空间复杂度O(n)
func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}

	return single
}
