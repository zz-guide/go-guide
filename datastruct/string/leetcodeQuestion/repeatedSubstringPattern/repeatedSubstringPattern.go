package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/repeated-substring-pattern/

重复的子字符串

给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

*/
func main() {
	s := "abcabcabcabc"
	fmt.Println("重复的子字符串:", repeatedSubstringPattern(s))
}

func repeatedSubstringPattern(s string) bool {

}
