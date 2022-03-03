package main

import (
	"log"
)

/**
题目：https://leetcode-cn.com/problems/find-the-winner-of-the-circular-game/

找出游戏的获胜者
共有 n 名小伙伴一起做游戏。小伙伴们围成一圈，按 顺时针顺序 从 1 到 n 编号。确切地说，从第 i 名小伙伴顺时针移动一位会到达第 (i+1) 名小伙伴的位置，其中 1 <= i < n ，从第 n 名小伙伴顺时针移动一位会回到第 1 名小伙伴的位置。

游戏遵循如下规则：

从第 1 名小伙伴所在位置 开始 。
沿着顺时针方向数 k 名小伙伴，计数时需要 包含 起始时的那位小伙伴。逐个绕圈进行计数，一些小伙伴可能会被数过不止一次。
你数到的最后一名小伙伴需要离开圈子，并视作输掉游戏。
如果圈子中仍然有不止一名小伙伴，从刚刚输掉的小伙伴的 顺时针下一位 小伙伴 开始，回到步骤 2 继续执行。
否则，圈子中最后一名小伙伴赢得游戏。
给你参与游戏的小伙伴总数 n ，和一个整数 k ，返回游戏的获胜者。

提示：

1 <= k <= n <= 500

*/
func main() {
	n := 5
	k := 2
	log.Println("约瑟夫环-动态规划:", findTheWinner(n, k))
	log.Println("约瑟夫环-模拟:", findTheWinner1(n, k))
	log.Println("约瑟夫环-递归:", findTheWinner2(n, k))
}

// findTheWinner 动态规划
func findTheWinner(n int, k int) int {
	//约瑟夫环，我吃定了
	//迭代解法
	dp := 0 //一个环的解就为最开始的小朋友的标号
	for i := 2; i <= n; i++ {
		dp = (dp + k) % i //dp为剩下的数字，而不是该删除的数字
	}
	return dp + 1 //下标和实际数字相差1
}

// findTheWinner1 模拟
func findTheWinner1(n int, k int) int {
	t := make([]int, n)
	for i := range t {
		t[i] = i + 1
	}

	cur := 0
	for i := 0; i < n-1; i++ {
		cur += k
		//fmt.Println(cur)
		cur %= len(t)
		//fmt.Println(t)
		if cur == 0 {
			t = t[:len(t)-1]
			continue
		}

		t = append(t[:cur-1], t[cur:]...)
		cur--
	}

	return t[0]
}

// findTheWinner2 递归
func findTheWinner2(n int, k int) int {
	if n == 1 {
		return 1
	}

	ans := findTheWinner(n-1, k) + k
	if ans%n == 0 {
		return n
	}

	return ans % n

}
