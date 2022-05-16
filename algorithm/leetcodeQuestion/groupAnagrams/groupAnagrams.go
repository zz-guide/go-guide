package main

import (
	"log"
	"sort"
)

/**
题目：https://leetcode-cn.com/problems/group-anagrams/solution/zi-mu-yi-wei-ci-fen-zu-by-leetcode-solut-gyoc/

字母异位词分组

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]

提示：

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i]仅包含小写字母

*/
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	log.Println("字母异位词分组-排序:", groupAnagrams(strs))
	log.Println("字母异位词分组-计数:", groupAnagrams1(strs))
}

// groupAnagrams 排序 O(nklogk) O(nk)
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

// groupAnagrams1 计数 O(n(k+∣Σ∣)) O(n(k+∣Σ∣))
func groupAnagrams1(strs []string) [][]string {
	// key可以是array
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}

	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
