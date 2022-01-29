package main

import (
	"log"
	"strconv"
)

/**
题目：https://leetcode-cn.com/problems/add-strings/solution/zi-fu-chuan-xiang-jia-by-leetcode-solution/

字符串相加
给定两个字符串形式的非负整数num1 和num2，计算它们的和并同样以字符串形式返回。

你不能使用任何內建的用于处理大整数的库（比如 BigInteger），也不能直接将输入的字符串转换为整数形式。

提示：

1 <= num1.length, num2.length <= 104
num1 和num2 都只包含数字0-9
num1 和num2 都不包含任何前导零


注意：1.字符串不能直接转数字

*/
func main() {
	num1 := "11"
	num2 := "123"
	log.Println("字符串相加:", addStrings(num1, num2))
}

// addStrings 模拟竖式加法，从最低位置加起，进位， O(max(len1,len2)）
func addStrings(num1 string, num2 string) string {
	jinwei := 0
	ans := ""

	// jinwei != 0 的时候还需要再重新执行一次，可能超位数
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || jinwei != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}

		if j >= 0 {
			y = int(num2[j] - '0')
		}

		// 计算当前位置的和
		result := x + y + jinwei
		// 重新计算进位
		jinwei = result / 10
		ans = strconv.Itoa(result%10) + ans
	}

	return ans
}
