package main

func main() {
	x()()
}

// 结果是：不停的输出zzzzz
func x() (y func()) {
	y = func() {
		println("yyyyy")
	}

	return func() {
		println("zzzzz")
		y()
	}
}
