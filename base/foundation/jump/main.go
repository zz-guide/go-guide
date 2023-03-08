package main

func main() {
	//TGoto()
	TBreak()
}

func TGoto() {
	// 结论：标签区分大小写，不能跳转到函数内部或内部代码块
	for i := 0; i < 3; i++ {
		if i > 1 {
			goto exit
		}

		println(i)
	}
exit:
	println("exit.")

}

func TBreak() {

	// break跳出多层循环可以使用标签法，或者goto,建议使用标签法
BBB:
	for i := 0; i < 2; i++ {
		if i%2 == 0 {
			continue BBB
		}

		println("i == ", i)

	AAA:
		for a := 0; a < 2; a++ {
			if a == 1 {
				break AAA
			}

			println("a == ", a)
		}
	}
}
