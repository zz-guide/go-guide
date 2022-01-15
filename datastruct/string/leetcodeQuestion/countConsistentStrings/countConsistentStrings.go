package main

/**
	题目：统计一致字符串的数目
	给你一个由不同字符组成的字符串allowed和一个字符串数组words。如果一个字符串的每一个字符都在 allowed中，就称这个字符串是 一致字符串 。

	请你返回words数组中一致字符串 的数目。
	示例 1：

	输入：allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
	输出：2
	解释：字符串 "aaab" 和 "baa" 都是一致字符串，因为它们只包含字符 'a' 和 'b' 。
	示例 2：

	输入：allowed = "abc", words = ["a","b","c","ab","ac","bc","abc"]
	输出：7
	解释：所有字符串都是一致的。
	示例 3：

	输入：allowed = "cad", words = ["cc","acd","b","ba","bac","bad","ac","d"]
	输出：4
	解释：字符串 "cc"，"acd"，"ac" 和 "d" 是一致字符串。

	提示：

	1 <= words.length <= 104
	1 <= allowed.length <= 26
	1 <= words[i].length <= 10
	allowed中的字符 互不相同。
	words[i] 和allowed只包含小写英文字母。
 */
func main() {

}

// F1
//**
func F1(allowed string, words []string) int {
	var allowedMap int
	for _, v := range allowed {
		allowedMap |= 1 << (v - 'a')
	}
	ans := 0
	var ok bool
	for _, v := range words {
		ok = true
		for _, c := range v {
			if allowedMap&(1<<(c-'a')) == 0 {
				ok = false
				break
			}
		}
		if ok {
			ans++
		}
	}
	return ans
}

// F2
//*
func F2(allowed string, words []string) int {
	allowedMap := [26]bool{}
	for _, v := range allowed {
		allowedMap[v-'a'] = true
	}
	ans := 0
	var ok bool
	for _, v := range words {
		ok = true
		for _, c := range v {
			if !allowedMap[c-'a'] {
				ok = false
				break
			}
		}
		if ok {
			ans++
		}
	}
	return ans
}

// F3
//**
func F3(allowed string, words []string) int {
	allowedMap := make(map[rune]bool)
	for _, v := range allowed {
		allowedMap[v] = true
	}
	ans := 0
	var ok bool
	for _, v := range words {
		ok = true
		for _, c := range v {
			if !allowedMap[c] {
				ok = false
				break
			}
		}
		if ok {
			ans++
		}
	}
	return ans
}
