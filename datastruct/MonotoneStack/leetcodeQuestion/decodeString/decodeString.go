package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/decode-string/
字符串解码

给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像3a或2[4]的输入。


注意：1.字符串没有空格等其他情况，正规
2.只要是数字就表示是次数
3.k是正整数，不只是0~9,且数字连续
4.单个字符是ASCII码，0是48

*/
func main() {
	s := "123"
	fmt.Println("字符串解码-递归:", decodeString(s))
	fmt.Println("字符串解码-单调栈:", decodeString1(s))
}

// decodeString 递归
func decodeString(s string) string {

	var decode func(start int) (string, int)
	decode = func(start int) (str string, end int) {
		num := 0
		for i := start; i < len(s); i++ {
			if isNumber(s[i]) {
				num = num*10 + int(s[i]-'0')
			} else if isLetter(s[i]) {
				str += string(s[i])
			} else if s[i] == '[' {
				item, index := decode(i + 1)
				for num != 0 {
					str += item
					num--
				}
				i = index
			} else if s[i] == ']' {
				end = i
				break
			}
		}

		return str, end
	}

	res, _ := decode(0)
	return res
}

// decodeString1 单调栈
func decodeString1(s string) string {
	n := len(s)

	var k int      // 次数，一个完整格式的次数
	var str string // 完整字符串

	numStack := make([]int, 0)
	strStack := make([]string, 0)

	for i := 0; i < n; i++ {
		if isNumber(s[i]) {
			k = k*10 + int(s[i]-'0')
		} else if isLetter(s[i]) {
			str += string(s[i])
		} else if s[i] == '[' {
			numStack = append(numStack, k)
			strStack = append(strStack, str) // 说明是新的一组开始，保存上一组的结果
			k, str = 0, ""
		} else if s[i] == ']' {
			repeatTimes, item := numStack[len(numStack)-1], strStack[len(strStack)-1]
			numStack, strStack = numStack[:len(numStack)-1], strStack[:len(strStack)-1]

			for j := 0; j < repeatTimes; j++ {
				item += str
			}

			str = item
		}
	}

	return str
}

func isLetter(u byte) bool {
	return u >= 'A' && u <= 'Z' || u >= 'a' && u <= 'z'
}

func isNumber(u byte) bool {
	return u >= '0' && u <= '9'
}
