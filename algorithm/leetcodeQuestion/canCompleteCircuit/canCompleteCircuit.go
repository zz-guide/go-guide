package main

import "log"

/**
题目：https://leetcode-cn.com/problems/gas-station/

加油站

在一条环路上有 n个加油站，其中第 i个加油站有汽油gas[i]升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1个加油站需要消耗汽油cost[i]升。你从其中的一个加油站出发，开始时油箱为空。

给定两个整数数组 gas 和 cost ，如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。

示例1:

输入: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
输出: 3
解释:
从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
因此，3 可为起始索引。
示例 2:

输入: gas = [2,3,4], cost = [3,4,3]
输出: -1
解释:
你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
因此，无论怎样，你都不可能绕环路行驶一周。

提示:

gas.length == n
cost.length == n
1 <= n <= 105
0 <= gas[i], cost[i] <= 104

*/
func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	log.Println("加油站:", canCompleteCircuit(gas, cost))
}

// canCompleteCircuit 贪心算法
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas) // len(gas) == len(cost) == n
	// 因为没有规定起点，所以一个一个开始作为起点尝试
	for i := 0; i < n; {
		sumGas, sumCost, cnt := 0, 0, 0 // 累计剩余油量，累计消耗油量，次数
		for cnt < n {
			j := (i + cnt) % n // 相对位置
			sumGas += gas[j]
			sumCost += cost[j]
			if sumGas < sumCost {
				break
			}

			cnt++
		}

		if cnt == n {
			return i
		} else {
			// i~i+cnt+1这段内，i不满足条件，那么区间内肯定不满足条件，
			i += cnt + 1
		}

	}

	return -1
}

// canCompleteCircuit2 贪心算法
func canCompleteCircuit2(gas []int, cost []int) int {
	start, totalSum, curSum := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		totalSum += gas[i] - cost[i]
		// 每次向下走的时候判断当前剩余
		curSum += gas[i] - cost[i]
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}

	if totalSum < 0 {
		return -1
	}

	return start
}
