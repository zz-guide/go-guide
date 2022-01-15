package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/happy-number/

编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果 可以变为 1，那么这个数就是快乐数。
如果 n 是快乐数就返回 true ；不是，则返回 false 。

提示：

1 <= n <= 231 - 1

*/
func main() {
	n := 19
	fmt.Println("快乐数-双指针:", isHappy(n))
	fmt.Println("快乐数-哈希:", isHappy1(n))
	fmt.Println("快乐数-数学归纳总结法:", isHappy2(n))
}

// isHappy 快慢指针，快指针每次执行2次，慢指针执行一次
func isHappy(n int) bool {
	slow, fast := n, getSum(n)
	for fast != 1 && slow != fast {
		slow = getSum(slow)
		fast = getSum(getSum(fast))
	}
	return fast == 1
}

// isHappy 哈希
func isHappy1(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}

	return n == 1
}

// isHappy2 数学归纳总结法
func isHappy2(n int) bool {
	cycle := map[int]bool{4: true, 6: true, 37: true, 58: true, 89: true, 145: true, 42: true, 20: true}
	for n != 1 && !cycle[n] {
		n = getSum(n)
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
