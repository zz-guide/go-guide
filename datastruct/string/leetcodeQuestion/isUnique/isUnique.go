package main

import "log"

/**
题目：https://leetcode-cn.com/problems/is-unique-lcci/solution/wei-yun-suan-by-ba-li-you-an-provenv-danz/

判定字符是否唯一

实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

示例 1：

输入: s = "leetcode"
输出: false
示例 2：

输入: s = "abc"
输出: true
限制：

0 <= len(s) <= 100
如果你不使用额外的数据结构，会很加分。

*/
func main() {
	s := "leetcode"
	log.Println("判定字符是否唯一:", isUnique(s))
}

// isUnique 位运算 O(1)空间，O(n)时间
// mark 长度为26，初始bit都是0
func isUnique(astr string) bool {
	if astr == "" {
		return true
	}

	mark := 0
	for _, r := range astr {
		moveBit := 1 << (r - 'a')
		// 说明已经存在
		if mark&moveBit != 0 {
			return false
		} else {
			// 或运算设置bit为1
			mark = mark | moveBit
		}
	}

	return true
}
