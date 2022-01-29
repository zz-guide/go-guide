package main

import (
	. "go-guide/datastruct/trie/trie"
	"log"
)

/**
题目：https://leetcode-cn.com/problems/implement-trie-prefix-tree/

Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。

请你实现 Trie 类：

Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

提示：

1 <= word.length, prefix.length <= 2000
word 和 prefix 仅由小写英文字母组成
insert、search 和 startsWith 调用次数 总计 不超过 3 * 104 次

使用场景:
字符串查找，词频统计，智能提示，敏感词过滤等。

*/
func main() {
	Do()
}

func Do() {
	words := []string{"insert"}
	trie := NewTrie()
	for _, word := range words {
		trie.Insert(word)
	}

	// 搜索
	res := trie.Search("inserta")
	log.Println("Search:", res)

	res1 := trie.StartsWith("in")
	log.Println("StartsWith:", res1)

	// 删除
	trie.Del("insert")
	res2 := trie.Search("insert")
	log.Println("Del之后再查找:", res2, trie.Children)
}
