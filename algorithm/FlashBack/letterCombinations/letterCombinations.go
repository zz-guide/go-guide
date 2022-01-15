package main

import "fmt"

/**

题目: https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
电话号码的字母组合

给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

提示：

0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。


注意：1.同组的字母之间不能相互组合
2.同组之间的字母只能单个成结果
3.有几个数字就需要几组for循环，此时应该使用回溯
4.组合必须每组取一个，不能不取,题目没有明确说，比较**

*/
func main() {
	digits := "2345"
	fmt.Println("电话号码的字母组合-回溯:", letterCombinations(digits))
}

func letterCombinations(digits string) []string {
	length := len(digits)
	if length == 0 || length > 4 {
		return nil
	}

	var digitsMap = [10]string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}

	var res []string
	var backtrack func(index int, combination string)
	backtrack = func(index int, track string) {
		if index == len(digits) {
			res = append(res, track)
			return
		}

		// 因为byte的关系，需要减去'0'
		letters := digitsMap[digits[index]-'0']
		for i := 0; i < len(letters); i++ {
			backtrack(index+1, track+string(letters[i]))
		}
	}

	backtrack(0, "")
	return res
}
