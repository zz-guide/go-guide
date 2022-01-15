package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/remove-k-digits/
移掉 K 位数字

给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。

输入：num = "1432219", k = 3
输出："1219"
解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219 。

提示：

1 <= k <= num.length <= 105
num 仅由若干位数字（0 - 9）组成
除了 0 本身之外，num 不含任何前导零


*/
func main() {
	num := "1432219"
	k := 3
	fmt.Println("移掉K位数字返回最小数字:", removeKdigits(num, k))
}

func removeKdigits(num string, k int) string {

}
