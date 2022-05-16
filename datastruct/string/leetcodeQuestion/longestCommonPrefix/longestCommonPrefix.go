package main

import "fmt"

/**
题目：
	编写一个函数来查找字符串数组中的最长公共前缀。

	如果不存在公共前缀，返回空字符串""。
	示例 1：

	输入：strs = ["flower","flow","flight"]
	输出："fl"
	示例 2：

	输入：strs = ["dog","racecar","car"]
	输出：""
	解释：输入不存在公共前缀。

	提示：

	1 <= strs.length <= 200
	0 <= strs[i].length <= 200
	strs[i] 仅由小写英文字母组成
*/
func main() {
	strs := []string{"abc", "ab", "abc", "abvd", "abcde"}
	prefix := F1(strs)
	fmt.Println("横向扫描：", prefix)
}

// F1 横向扫描
//时间复杂度：O(mn)O(mn)，其中 mm 是字符串数组中的字符串的平均长度，nn 是字符串的数量。最坏情况下，字符串数组中的每个字符串的每个字符都会被比较一次。
//空间复杂度：O(1)O(1)。使用的额外空间复杂度为常数
//
//**
func F1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	// 思路：给定2个字符串，谁的长度小，遍历谁，通过索引取到每个字符
	x := len(str1)
	y := len(str2)
	length := 0
	if x < y {
		length = x
	} else {
		length = y
	}

	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}

	if index == 0 {
		return ""
	}

	return str1[:index] // 返回str1或者str2都行
}

// F2 纵向扫描
// 时间复杂度：O(mn)O(mn)，其中 mm 是字符串数组中的字符串的平均长度，nn 是字符串的数量。最坏情况下，字符串数组中的每个字符串的每个字符都会被比较一次。
// 空间复杂度：O(1)O(1)。使用的额外空间复杂度为常数。
//**
func F2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// F3
//时间复杂度：O(mn)O(mn)，其中 mm 是字符串数组中的字符串的平均长度，nn 是字符串的数量。时间复杂度的递推式是 T(n)=2 \cdot T(\frac{n}{2})+O(m)T(n)=2⋅T(
//2
//n
//)+O(m)，通过计算可得 T(n)=O(mn)T(n)=O(mn)。
//空间复杂度：O(m \log n)O(mlogn)，其中 mm 是字符串数组中的字符串的平均长度，nn 是字符串的数量。空间复杂度主要取决于递归调用的层数，层数最大为 \log nlogn，每层需要 mm 的空间存储返回结果。
//**
func F3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2

		lcpLeft, lcpRight := lcp(start, mid), lcp(mid+1, end)

		x := len(lcpLeft)
		y := len(lcpRight)
		minLength := 0
		if x < y {
			minLength = x
		} else {
			minLength = y
		}

		for i := 0; i < minLength; i++ {
			if lcpLeft[i] != lcpRight[i] {
				return lcpLeft[:i]
			}
		}
		return lcpLeft[:minLength]
	}
	return lcp(0, len(strs)-1)
}

//F4 二分查找
//时间复杂度：O(mn \log m)O(mnlogm)，其中 mm 是字符串数组中的字符串的最小长度，nn 是字符串的数量。二分查找的迭代执行次数是 O(\log m)O(logm)，每次迭代最多需要比较 mnmn 个字符，因此总时间复杂度是 O(mn \log m)O(mnlogm)。
//空间复杂度：O(1)O(1)。使用的额外空间复杂度为常数
//**
func F4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	isCommonPrefix := func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}

	minLength := len(strs[0])
	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}

	low, high := 0, minLength
	for low < high {
		mid := (high-low+1)/2 + low
		if isCommonPrefix(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}

	return strs[0][:low]
}
