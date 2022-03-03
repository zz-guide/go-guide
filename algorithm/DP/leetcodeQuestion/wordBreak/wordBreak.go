package main

import "log"

/**
题目：https://leetcode-cn.com/problems/word-break/

单词拆分

给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。

注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
   注意，你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false

提示：

1 <= s.length <= 300
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 20
s 和 wordDict[i] 仅有小写英文字母组成
wordDict 中的所有字符串 互不相同


*/
func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	log.Println("单词拆分-动态规划:", wordBreak(s, wordDict))
	log.Println("单词拆分-记忆化回溯:", wordBreak2(s, wordDict))
	log.Println("单词拆分-前缀树:", wordBreak3(s, wordDict))
}

// wordBreak 动态规划 O(n^2)
func wordBreak(s string, wordDict []string) bool {
	// 构建map
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}

	dp := make([]bool, len(s)+1) // dp[i] 表示字符串 ss 前 ii 个字符组成的字符串 s[0..i-1]s[0..i−1] 是否能被空格拆分成若干个字典中出现的单词
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			// dp[j]+s[j:i]=dp[i]
			if wordMap[s[j:i]] && dp[j] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

// wordBreak2 回溯+记忆化
func wordBreak2(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	// 避免每次都要去slice里遍历匹配
	for _, v := range wordDict {
		wordMap[v] = true
	}

	cache := make(map[int]bool)
	var dfs func(start int, s string, cache map[int]bool) bool
	dfs = func(start int, s string, cache map[int]bool) bool {
		if v, ok := cache[start]; ok {
			return v
		}

		if start == len(s) {
			return true
		}

		for i := start + 1; i < len(s)+1; i++ {
			prefix := s[start:i]
			// 当前前缀是字典里的单词并且剩余的字符串也是字典里的单词
			if wordMap[prefix] && dfs(i, s, cache) {
				// 避免重复计算
				// 从这个start位置开始的字符串是字典里的单子
				cache[start] = true
				return true
			}
		}
		// 当前层没有纵向和横向返回的结果都不是，从这个index开始的字符串不是字典里的单词
		cache[start] = false
		// 返回给上层
		return false
	}

	return dfs(0, s, cache)
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Trie) dfs(word string, startPos int, failMemo []bool) bool {
	if startPos == len(word) {
		return true
	}
	if failMemo[startPos] == true {
		return false
	}
	node := this
	for i := startPos; i < len(word); i++ {
		node = node.children[word[i]-'a']
		if node == nil {
			break
		}
		if node.isEnd && this.dfs(word, i+1, failMemo) {
			return true
		}
	}
	failMemo[startPos] = true
	return false
}

func wordBreak3(s string, wordDict []string) bool {
	trie := &Trie{}
	for _, word := range wordDict {
		trie.Insert(word)
	}
	return trie.dfs(s, 0, make([]bool, len(s)))
}
