package main

import (
	"fmt"
	"strings"
)

/**
题目：
	给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
	说明：本题中，我们将空字符串定义为有效的回文串。
	示例 1:

	输入: "A man, a plan, a canal: Panama"
	输出: true
	解释："amanaplanacanalpanama" 是回文串
	示例 2:

	输入: "race a car"
	输出: false
	解释："raceacar" 不是回文串

	思路：
		1.回文串特点：反转之后和反转前的字符串相等
		2.特殊的字符串去掉特殊字符后符合第一条规则
		3.本质就是比较对应位置的字符是否相等（i和length1-i）
		4.遍历+双指针或者遍历交换位置
*/
func main() {
	str := "A man, a plan, a canal: Panama"
	result := F1(str)
	fmt.Println("是回文串？", result)
}

// F1 反转字符串函数实现，交换对应位置的值即可
//	因为有其他符号所以先过滤掉，因为大小写不敏感，所以统一转成小写
//**
func F1(str string) bool {
	var sgood string
	for i := 0; i < len(str); i++ {
		ch := str[i]
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			sgood += string(str[i])
		}
	}

	str = strings.ToLower(sgood)
	strArr := []rune(str)
	length := len(strArr)
	for i := 0; i < length/2; i++ {
		strArr[length-1-i], strArr[i] = strArr[i], strArr[length-1-i]
	}

	return string(strArr) == str
}

// F2 双指针写法
//时间复杂度：O(|s6|)O(∣s6∣)，其中 |s6|∣s6∣ 是字符串 ss 的长度。
//空间复杂度：O(|s6|)O(∣s6∣)。由于我们需要将所有的字母和数字字符存放在另一个字符串中，在最坏情况下，新的字符串 \textit{sgood}sgood 与原字符串 ss 完全相同，因此需要使用 O(|s6|)O(∣s6∣) 的空间
//**
func F2(s string) bool {
	var sgood string
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			sgood += string(s[i])
		}
	}

	length := len(sgood)
	sgood = strings.ToLower(sgood)
	for i := 0; i < length/2; i++ {
		if sgood[i] != sgood[length-1-i] {
			return false
		}
	}
	return true
}

// F3 直接在原来的字符串上进行判断，优化空间复杂度，本质上也是双指针
//
func F3(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}

		for left < right && !isalnum(s[right]) {
			right--
		}

		if left < right {
			if s[left] != s[right] {
				return false
			}

			left++
			right--
		}
	}

	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
