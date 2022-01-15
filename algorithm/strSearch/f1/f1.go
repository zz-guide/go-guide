package f1

// StrStr
// 暴力匹配:时间复杂度：O(n×m) 空间复杂度：O(1)
///**
func StrStr(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if n < m {
		return -1
	}

	if len(needle) == 0 { //若模式串为空串
		return 0
	}

outer:
	for i := 0; i <= n-m; i++ {
		for j := range needle {
			if haystack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}
