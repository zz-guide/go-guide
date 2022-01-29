package main

import (
	"log"
)

/**
什么是KMP
因为是由这三位学者发明的：Knuth，Morris和Pratt，所以取了三位学者名字的首字母。所以叫做KMP

KMP有什么用
KMP主要应用在字符串匹配上。主要思想是当出现字符串不匹配时，可以知道一部分之前已经匹配的文本内容，可以利用这些信息避免从头再去做匹配了。

什么是前缀表
写过KMP的同学，一定都写过next数组，那么这个next数组究竟是个啥呢？

next数组就是一个前缀表（prefix table）。
前缀表是用来回退的，它记录了模式串与主串(文本串)不匹配的时候，模式串应该从哪里开始重新匹配。

KMP的时间复杂度是：匹配过程的时间复杂度为O(n)，计算next的O(m)时间，两个独立的环节串行，所以整体时间复杂度为O(m + n)。

https://blog.csdn.net/sunnianzhong/article/details/8802559

*/

func main() {
	str := "ababd"
	target := "abd"
	log.Println("KMP算法查找:", KMP(str, target))
}

// KMP 构造前缀表next 方法一:前缀表使用减1实现
// params: next 前缀表数组 s6 模式串
func KMP(str string, target string) int {
	if len(target) == 0 {
		return 0
	}

	if len(target) > len(str) {
		return -1
	}

	next := getNext(target)

	j := 0
	for i := 0; i < len(str); i++ {
		for j > 0 && str[i] != target[j] {
			j = next[j-1]
		}

		if str[i] == target[j] {
			j++
		}

		if j == len(target) {
			return i - len(target) + 1
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
