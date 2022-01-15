package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/validate-stack-sequences/

给定pushed和popped两个序列，每个序列中的 值都不重复，只有当它们可能是在最初空栈上进行的推入 push 和弹出 pop 操作序列的结果时，返回 true；否则，返回 false。

验证栈序列

提示：

1 <= pushed.length <= 1000
0 <= pushed[i] <= 1000
pushed 的所有元素 互不相同
popped.length == pushed.length
popped 是 pushed 的一个排列

注意：1.两个序列没有重复的值
2.题目的意思是先push,当push的元素与pop的第一个元素相同时，同时出栈，继续比较，知道pop最后变成0，就是true

*/
func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	fmt.Println("验证栈序列:", validateStackSequences(pushed, popped))
}

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	var res bool
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) != 0 && popped[0] == stack[len(stack)-1] {
			popped = popped[1:]
			stack = stack[:len(stack)-1]
		}
	}

	if len(popped) == 0 {
		res = true
	}

	return res
}
