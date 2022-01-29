package main

import (
	"log"
	"strconv"
)

/**
题目：https://leetcode-cn.com/problems/restore-ip-addresses/
复原 IP 地址

有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入'.' 来形成。你不能重新排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。

提示：

0 <= s.length <= 20
s 仅由数字组成


注意：1.s中插入.让s变成ip地址串，不能删除，也不能重新排序
2.类似分隔回文串题目，只需很少改动
3.4个一组，用.隔开

*/
func main() {
	s := "25525511135"
	log.Println("复原 IP 地址:", restoreIpAddresses(s))
}

func restoreIpAddresses(s string) []string {
	length := len(s)
	var res []string
	var isIp func(l, r int) bool
	isIp = func(l, r int) bool {
		checkInt, _ := strconv.Atoi(s[l : r+1])
		// 判断前导0
		if r-l+1 > 1 && s[l] == '0' {
			return false
		}

		return checkInt >= 0 && checkInt <= 255
	}

	var track []string
	var backtracking func(start int)
	backtracking = func(start int) {
		if start == length && len(track) == 4 {
			res = append(res, track[0]+"."+track[1]+"."+track[2]+"."+track[3])
			return
		}

		for i := start; i < length; i++ {
			if isIp(start, i) {
				track = append(track, s[start:i+1])
				backtracking(i + 1)
				track = track[:len(track)-1]
			}
		}
	}

	backtracking(0)
	return res
}
