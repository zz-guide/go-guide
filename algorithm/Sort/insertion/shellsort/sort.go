package main

import "fmt"

/**
希尔排序(Shell's6 Sort)是插入排序的一种又称“缩小增量排序”（Diminishing Increment Sort），是直接插入排序算法的一种更高效的改进版本。
希尔排序是非稳定排序算法。该方法因 D.L.Shell 于 1959 年提出而得名。
希尔排序是把记录按下标的一定增量分组，对每组使用直接插入排序算法排序；随着增量逐渐减少，每组包含的关键词越来越多，当增量减至 1 时，整个文件恰被分成一组，算法便终止。
时间复杂度：O(n^（1.3—2）)
空间复杂度：O(1)
稳定性：不稳定
*/
func main() {
	list := []int{4, 3, 5, 1, 2, 9}
	fmt.Println("希尔排序:", ShellSort(list))
}

func ShellSort(nums []int) []int {
	//外层步长控制
	for step := len(nums) / 2; step > 0; step /= 2 {
		//开始插入排序
		for i := step; i < len(nums); i++ {
			//满足条件则插入
			for j := i - step; j >= 0 && nums[j+step] < nums[j]; j -= step {
				nums[j], nums[j+step] = nums[j+step], nums[j]
			}
		}
	}
	return nums
}
