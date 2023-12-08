package fib

import "testing"

/**
	https://geektutu.com/post/hpg-empty-struct.html
	testing 支持生成 CPU、memory 和 block 的 profile 文件。

	-cpuprofile=$FILE
	-memprofile=$FILE, -m
	-blockprofile=$FILE

	go test -bench="fib$" -cpuprofile=cpu.pprof .
	go tool pprof -http=:8888 ./cpu.pprof
 */

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}
}