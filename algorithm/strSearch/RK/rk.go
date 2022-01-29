package main

import (
	"fmt"
	"strings"
)

/**
RK算法全程Rabin-Karp，该算法的2位发明者Rabin和Karp的名字组合而成。该算法的核心思想就是通过比较2个字符串的hashcode来判断是否包含对方。

由于该算法核心思想是计算字符串的hashcode，因此必须保证hash算法针对不同的字符串得出不同的值，例如：abc、bca、acb这3种相同字符不同排序的情况下，保证这3种字符串的hashcode不同。

RK算法也可以进行多模式匹配，在论文查重等实际应用中一般都是使用此算法。

时间复杂度：O（MN）（实际应用中往往较快，期望时间为O（M+N））

使用HASH比较。
如果两个字符串hash后的值不相同，则它们肯定不相同；如果它们hash后的值相同，它们不一定相同。
RK算法的基本思想就是：将模式串P的hash值跟主串T中的每一个长度为|P|的子串的hash值比较。如果不同，则它们肯定不相等；如果相同，由于哈希冲突存在，也需要按照BF算法诸位比较。

Rabin-Karp算法思想：

假设待匹配字符串长度M，目标字符串长度N（N>M）
首先计算待匹配字符串hash，计算目标字符串前M个字符hash
比较前两个hash值，比较次数N-M+1
若hash不相等，继续计算目标字符串下一个长度为M的hash并继续循环比较
若hash相等则再次判断字符串是否相等已确保正确性

*/
func main() {
	str := "ababd"
	target := "abd"
	fmt.Println("RabinKarp算法查找:", IndexRabinKarp(str, target))
	strings.Index(str, target)
}

// Rabin-Karp 中需要使用的32位FNV hash算法中的基础质数（相当于进制）
const PrimeRK = 16777619

// hash散列方法， 返回字符串hash以及 primeRK的k-1（len(sep)-1）次方
// HashStr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func HashStr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*PrimeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, PrimeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

// IndexRabinKarp uses the Rabin-Karp search algorithm to return the index of the
// first occurrence of substr in s, or -1 if not present.
func IndexRabinKarp(s, substr string) int {
	// Rabin-Karp search
	hashss, pow := HashStr(substr)
	n := len(substr)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*PrimeRK + uint32(s[i])
	}
	if h == hashss && s[:n] == substr {
		return 0
	}
	for i := n; i < len(s); {
		h *= PrimeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashss && s[i-n:i] == substr {
			return i - n
		}
	}
	return -1
}
