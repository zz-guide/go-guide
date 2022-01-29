package main

import (
	"log"
)

/**
给你一个可装载重量为 W 的背包和 N 个物品，每个物品有重量和价值两个属性。其中第 i 个物品的重量为 wt[i]，价值为 val[i]，现在让你用这个背包装物品，最多能装的价值是多少？


定义一个函数 f(K, W) = val,N代表装几件物品，，V代表当前背包的体积，val代表最大价值

假设有N=4个物品，背包体积为W=8
索引   	物品   	体积    	价值
0     	物品1   	2      	3

1		物品2   	3      	4

2		物品3   	4       5

3		物品4   	5       8

每一件物品都有装和不装两种情况，当这样穷举每一件物品装和不装的最大价值情况就能得出结果


状态转移方程：
f(K,W)= [
	f(K-1, W), 不装,因为装不下, Wk > W
	max{f(K-1, W),  f(K-1, W - Wk)+ Vk(假设能装下)}, WK <= W
]

*/
func main() {
	goodsWeight := []int{2, 3, 4, 5}
	goodsValue := []int{3, 4, 5, 8}
	bagWeight := 8
	log.Println("背包最大价值(二维简单):", bagMaxValue1(goodsWeight, goodsValue, bagWeight))
	log.Println("背包最大价值(二维优化):", bagMaxValue2(goodsWeight, goodsValue, bagWeight))
	log.Println("背包最大价值(一维优化):", bagMaxValue3(goodsWeight, goodsValue, bagWeight))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// bagMaxValue1 二维数组常规解法，dp[N+1][C+1]
func bagMaxValue1(goodsWeight []int, value []int, bagWeight int) int {
	if bagWeight <= 0 || len(goodsWeight) != len(value) {
		return 0
	}

	// 定义dp数组，多分配一个，之后都从1开始计算
	// dp[N][W]含义：装N件物品，背包体积为W时的最大值
	dp := make([][]int, len(goodsWeight)+1)
	for i, _ := range dp {
		// 背包也多分配一个，从第一个背包开始计算
		dp[i] = make([]int, bagWeight+1)
	}

	// 假设从第一个物品开始装，因为多分配了一个位置，所以可以=len(goodsWeight)
	for i := 1; i <= len(goodsWeight); i++ {
		// 假设背包体积为1开始装，从大到小也可以
		for w := 1; w <= bagWeight; w++ {
			// 因为索引都从1开始，所以goodsWeight， value注意不要越界
			// 从状态方程的角度来写逻辑
			if w < goodsWeight[i-1] { // 若背包体积小于物品体积，装不下
				dp[i][w] = dp[i-1][w] // f(K-1, W),对于第一件物品没有之前的物品了
			} else {
				dp[i][w] = max(
					dp[i-1][w],                             // 不装前i-1件物品的最大值
					dp[i-1][w-goodsWeight[i-1]]+value[i-1], // 装前i-1件物品的最大值
				)
			}
		}
	}

	return dp[len(goodsWeight)][bagWeight]
}

// bagMaxValue2 二维数组优化解法，dp[2][C+1]
func bagMaxValue2(goodsWeight []int, value []int, bagWeight int) int {
	if bagWeight <= 0 || len(goodsWeight) != len(value) {
		return 0
	}

	// dp数组含义：[2][bagWeight+1],只需要保存当前物品的最大值和前一个物品的最大值就行，不需要存其他物品信息
	dp := make([][]int, 2)
	for i, _ := range dp {
		// 背包也多分配一个，从第一个背包开始计算
		dp[i] = make([]int, bagWeight+1)
	}

	// 假设从第一个物品开始装，因为多分配了一个位置，所以可以=len(goodsWeight)
	// i&1替代i/2
	for i := 1; i <= len(goodsWeight); i++ {
		for w := 1; w <= bagWeight; w++ {
			if w < goodsWeight[i-1] { // 若背包体积小于物品体积，装不下
				dp[i&1][w] = dp[(i-1)&1][w]
			} else {
				dp[i&1][w] = max(
					dp[(i-1)&1][w], // 不装前i-1件物品的最大值
					dp[(i-1)&1][w-goodsWeight[i-1]]+value[i-1], // 装前i-1件物品的最大值
				)
			}
		}
	}

	return dp[len(goodsWeight)&1][bagWeight]
}

// bagMaxValue3 一维数组解法，滚动数组(重点)
func bagMaxValue3(weight []int, value []int, maxWeight int) int {
	if maxWeight <= 0 || len(weight) != len(value) {
		return 0
	}

	// 多分配一个，M代表背包的体积对应的最大值
	M := make([]int, maxWeight+1)
	// 获取最小的物品体积
	min := weight[0]
	for _, w := range weight {
		if w < min {
			min = w
		}
	}

	for i := 0; i < len(weight); i++ {
		// 只遍历背包体积~最小物品体积区间范围
		for w := maxWeight; w >= min; w-- {
			if weight[i] <= w {
				M[w] = max(M[w], M[w-weight[i]]+value[i])
			}
		}
	}

	return M[maxWeight]
}
