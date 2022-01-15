package main

import "fmt"

/**
BF（Brute Force）算法
1.从主串的起始位置开始与子串逐个字符比较，遇到不相等的字符，改变主串的起始位置，从主串第二个字符开始比较，依次进行
2.很暴力，起始就是2层循环，外层主串，内层子串
3.O(m*n) O(1)

*/
func main() {
	str := "ababd"
	target := "abd"
	fmt.Println("BF算法查找:", BF(str, target))
}

// BF 外层遍历主串，内层遍历子串
func BF(str, target string) int {
	if len(target) == 0 {
		return -1
	}

	if len(str) == 0 || len(str) < len(target) {
		return -1
	}

	mStr := []rune(str)
	subStr := []rune(target)

	for i := 0; i < len(mStr); i++ {

		startIndex := i
		for j := 0; j < len(subStr); j++ {
			if mStr[startIndex] != subStr[j] {
				break
			}

			// 说明完全匹配了
			if j == len(subStr)-1 {
				return i
			}

			startIndex++
		}
	}

	return -1
}
