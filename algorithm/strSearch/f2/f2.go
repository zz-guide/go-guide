package f2

func StrStr(haystack string, needle string) int {
	if len(haystack) == 0 || len(haystack) < len(needle) {
		return -1
	}

	if len(needle) == 0 { //若模式串为空串
		return 0
	}

	M := len(haystack) //主串的长度
	N := len(needle)   //模式串的长度

	for i := 0; i <= M-N; i++ {
		if haystack[i:i+N] == needle {
			return i
		}
	}

	return -1
}
