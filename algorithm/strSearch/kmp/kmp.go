package kmp

/**
什么是KMP
说到KMP，先说一下KMP这个名字是怎么来的，为什么叫做KMP呢。

因为是由这三位学者发明的：Knuth，Morris和Pratt，所以取了三位学者名字的首字母。所以叫做KMP

KMP有什么用
KMP主要应用在字符串匹配上。主要思想是当出现字符串不匹配时，可以知道一部分之前已经匹配的文本内容，可以利用这些信息避免从头再去做匹配了。

什么是前缀表
写过KMP的同学，一定都写过next数组，那么这个next数组究竟是个啥呢？

next数组就是一个前缀表（prefix table）。
前缀表是用来回退的，它记录了模式串与主串(文本串)不匹配的时候，模式串应该从哪里开始重新匹配。

*/

// StrStr 构造前缀表next 方法一:前缀表使用减1实现
// params: next 前缀表数组 s6 模式串
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	next := getNext(needle)

	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}

		if haystack[i] == needle[j] {
			j++
		}

		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}

func getNext(s string) []int {
	next := make([]int, len(s))
	j := 0      // j表示 最长相等前后缀长度
	next[0] = j // 子串长度为1的时候相等长度为0

	for i := 1; i < len(s); i++ {

		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}

		// i和j相等的情况，长度等于j++
		if s[i] == s[j] {
			j++
		}

		// 当不相等的情况直接j=0
		next[i] = j
	}

	return next
}
