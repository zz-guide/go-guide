package firstuniqchar

import "fmt"

func firstUniqChar(s string) int {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}

	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return i
		}
	}

	return -1
}

func firstUniqChar1(s string) int {
	n := len(s)
	pos := [26]int{}
	for i := range pos[:] {
		pos[i] = n
	}
	for i, ch := range s {
		ch -= 'a'
		if pos[ch] == n {
			pos[ch] = i
		} else {
			pos[ch] = n + 1
		}
	}
	ans := n
	for _, p := range pos[:] {
		if p < ans {
			ans = p
		}
	}
	if ans < n {
		return ans
	}
	return -1
}

type pair struct {
	ch  byte
	pos int
}

func firstUniqChar3(s string) int {
	n := len(s)
	pos := [26]int{}
	for i := range pos[:] {
		pos[i] = n
	}
	var q []pair
	for i := range s {
		ch := s[i] - 'a'
		if pos[ch] == n {
			pos[ch] = i
			q = append(q, pair{ch, i})
		} else {
			pos[ch] = n + 1
			for len(q) > 0 && pos[q[0].ch] == n+1 {
				q = q[1:]
			}
		}
	}
	if len(q) > 0 {
		return q[0].pos
	}
	return -1
}

func DoTestFirstUniqChar() {
	s := "leetcode"
	i := firstUniqChar(s)
	fmt.Println("结果：", i)

	fmt.Println("---:", 'z', 'a')
}
