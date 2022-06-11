package main

import (
	"log"
	"strconv"
)

/**
字符串相乘
题目：https://leetcode.cn/problems/multiply-strings/
给定两个以字符串形式表示的非负整数num1和num2，返回num1和num2的乘积，它们的乘积也表示为字符串形式。

注意：不能使用任何内置的 BigInteger 库或直接将输入转换为整数。


示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例2:

输入: num1 = "123", num2 = "456"
输出: "56088"

提示：

1 <= num1.length, num2.length <= 200
num1和 num2只能由数字组成。
num1和 num2都不包含任何前导零，除了数字0本身。

*/
func main() {
	num1 := "123"
	num2 := "456"
	log.Println("字符串相乘-乘法:", multiply(num1, num2))
}

// 乘法 O(mn),O(m+n)O(m+n)
func multiply(num1 string, num2 string) string {
	// 只要任意一方为0，直接返回"0"
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	ansArr := make([]int, m+n) // 乘积的长度只可能是m+n或者m+n+1
	// 倒着进行乘法运算, (m-1)*(n-1)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i]) - '0'
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			// TODO:: 待确定为什么是i+j+1
			ansArr[i+j+1] += x * y
		}
	}

	log.Println("ansArr:", ansArr)

	// 这一步处理进位
	for i := m + n - 1; i > 0; i-- {
		// 前一位置的数+进位
		ansArr[i-1] += ansArr[i] / 10
		// 当前位置保留个位数
		ansArr[i] %= 10
	}

	res := ""
	for idx, v := range ansArr {
		if idx == 0 && v == 0 {
			continue
		}

		res += strconv.Itoa(v)
	}

	return res
}
