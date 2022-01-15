package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/daily-temperatures/

每日温度
请根据每日 气温 列表 temperatures，请计算在每一天需要等几天才会有更高的温度。如果气温在这之后都不会升高，请在该位置用0 来代替。

示例 1:
输入: temperatures = [73,74,75,71,69,72,76,73]
输出:[1,1,4,2,1,1,0,0]

示例 2:
输入: temperatures = [30,40,50,60]
输出:[1,1,1,0]

示例 3:
输入: temperatures = [30,60,90]
输出: [1,1,0]

提示：

1 <=temperatures.length <= 105
30 <=temperatures[i]<= 100

注意：
1.最后一个数字的结果肯定是0
2.假设j>i, t[j] > t[i],那么结果就是 j - i

*/
func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println("每日温度-暴力法:", dailyTemperatures(temperatures))
	fmt.Println("每日温度-单调栈:", dailyTemperatures1(temperatures))
}

// dailyTemperatures 暴力法
func dailyTemperatures(temperatures []int) []int {
	var res []int
	for i := 0; i < len(temperatures)-1; i++ {
		j := i + 1
		for ; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				res = append(res, j-i)
				break
			}
		}

		if j == len(temperatures) {
			res = append(res, 0)
		}
	}

	return append(res, 0)
}

//输入: temperatures = [73,74,75,71,69,72,76,73]
//输出:[1,1,4,2,1,1,0,0]
// dailyTemperatures1 单调递减栈
func dailyTemperatures1(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack []int // 假设是递减的
	for i, v := range temperatures {
		for len(stack) != 0 && v > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res[top] = i - top
		}

		stack = append(stack, i)
	}

	return res
}
