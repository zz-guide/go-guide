package strperformance

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/**
1. bytes.Buffer,strings.Builder,+,fmt.Sprintf,strings.Join
2. 大字符串推荐使用strings.Builder，底层没有拷贝数据，使用指针直接转换
*/

const numbers = 10000

func BenchmarkStringsJoin(b *testing.B) {
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		var sArr []string
		var s string
		for i := 0; i < numbers; i++ {
			sArr = append(sArr, s, strconv.Itoa(i))
			s = strings.Join(sArr, "")
		}
	}
	b.StopTimer()
}

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		for i := 0; i < numbers; i++ {
			var s string
			s = fmt.Sprintf("%v%v", s, i)
		}
	}
	b.StopTimer()
}

func BenchmarkStringsBuilder(b *testing.B) {
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		var builder strings.Builder
		for i := 0; i < numbers; i++ {
			builder.WriteString(strconv.Itoa(i))
		}
	}
	b.StopTimer()
}

func BenchmarkBytesBuffer(b *testing.B) {
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		var buf bytes.Buffer
		for i := 0; i < numbers; i++ {
			buf.WriteString(strconv.Itoa(i))
		}
	}
	b.StopTimer()
}

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		var s string
		for i := 0; i < numbers; i++ {
			s += strconv.Itoa(i)
		}
	}
	b.StopTimer()
}
