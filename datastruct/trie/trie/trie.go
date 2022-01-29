package trie

/**
前缀树或者字典树
Trie是一颗非典型的多叉树模型，多叉好理解，即每个结点的分支数量可能为多个。
我们可以看到TrieNode结点中并没有直接保存字符值的数据成员，那它是怎么保存字符的呢？

这时字母映射表next 的妙用就体现了，TrieNode* next[26]中保存了对当前结点而言下一个可能出现的所有字符的链接，因此我们可以通过一个父结点来预知它所有子结点的值：
Trie 中一般都含有大量的空链接，因此在绘制一棵单词查找树时一般会忽略空链接

一次建树，多次查询
以空间换时间，主要场景是用来做查询的

时间复杂度：初始化为 O(1)O(1)，其余操作为 O(|S|)，其中 |S|∣S∣ 是每次插入或查询的字符串的长度。

空间复杂度：O(|T|\cdot\Sigma)O(∣T∣⋅Σ)，其中 ∣T∣ 为所有插入字符串的长度之和，\SigmaΣ 为字符集的大小，本题 \Sigma=26Σ=26

*/

type Trie struct {
	Children [26]*Trie // 只允许小写字母，字符 - 'a'就是index,使用map也可以
	IsEnd    bool
	Val      rune
	Depth    int
	Count    int // 统计分支数量
}

func Constructor() Trie {
	return NewTrie()
}

func NewTrie() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.Children[ch] == nil {
			node.Children[ch] = &Trie{
				Val:   ch + 'a',
				Depth: node.Depth + 1,
			}

			node.Count += 1
		}
		node = node.Children[ch]
	}
	node.IsEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.Children[ch] == nil {
			return nil
		}
		node = node.Children[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.IsEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

func (t *Trie) Del(word string) {
	if len(word) == 0 {
		return
	}

	node := t
	var lastBranch *Trie
	var lastBranchIndex int
	for index, ch := range word {
		ch -= 'a'
		if node.Children[ch] == nil {
			break
		}

		if node.Children[ch].Count > 1 {
			lastBranch = node.Children[ch]
			lastBranchIndex = index
		}

		node = node.Children[ch]
	}

	// 有分支
	if node.Count > 0 {
		node.IsEnd = false
		return
	}

	// 没有分支
	if lastBranch == nil {
		// 删除整个字符串
		lastBranch = t
	}

	// 只能删除这一个字符
	// 其实应该递归的删除Count只有1个的字符。不然空间也不会减少，没有意义，跟bloom一样
	lastBranch.Children[lastBranchIndex] = nil
	lastBranch.Count -= 1
	return
}
