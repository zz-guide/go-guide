package main

import "fmt"

/**
给你一个字符串 s 和一个 长度相同 的整数数组 indices 。
请你重新排列字符串 s ，其中第 i 个字符需要移动到 indices[i] 指示的位置。
返回重新排列后的字符串。
*/

func main() {

}

func restoreString(s string, indices []int) string {
	tempArr := make([]rune, len(s))
	for index, value := range indices {
		tempArr[value] = rune(s[index])
	}

	return string(tempArr)
}

func DoRestoreString() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	result := restoreString(s, indices)
	fmt.Println("result:", result)
}
