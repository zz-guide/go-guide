package main

import (
	"log"
	"strings"
)

/**
题目：https://leetcode-cn.com/problems

单词拆分 II

给定一个字符串 s 和一个字符串字典 wordDict ，在字符串 s 中增加空格来构建一个句子，使得句子中所有的单词都在词典中。以任意顺序 返回所有这些可能的句子。

注意：词典中的同一个单词可能在分段中被重复使用多次。

示例 1：

输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
输出:["cats and dog","cat sand dog"]
示例 2：

输入:s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine","pineapple"]
输出:["pine apple pen apple","pineapple pen apple","pine applepen apple"]
解释: 注意你可以重复使用字典中的单词。
示例3：

输入:s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
输出:[]

提示：

1 <= s.length <= 20
1 <= wordDict.length <= 1000
1 <= wordDict[i].length <= 10
s和wordDict[i]仅有小写英文字母组成
wordDict中所有字符串都 不同

*/
func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	log.Println("单词拆分-回溯:", wordBreak(s, wordDict))
}

// wordBreak 回溯
func wordBreak(s string, wordDict []string) (sentences []string) {
	wordSet := map[string]struct{}{}
	for _, w := range wordDict {
		wordSet[w] = struct{}{}
	}

	n := len(s)
	dp := make([][][]string, n)

	var backtrack func(index int) [][]string
	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}

		var wordsList [][]string
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if _, has := wordSet[word]; has {
				for _, nextWords := range backtrack(i) {
					wordsList = append(wordsList, append([]string{word}, nextWords...))
				}
			}
		}

		word := s[index:]
		if _, has := wordSet[word]; has {
			wordsList = append(wordsList, []string{word})
		}

		dp[index] = wordsList
		return wordsList
	}

	for _, words := range backtrack(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}

	return
}
