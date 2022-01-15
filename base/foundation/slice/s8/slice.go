package slice

import "fmt"

/**
slice（切片）是一种动态数组，其底层为Array,连续分配
容量与长度区别：容量就是最大长度
*/
func TestSliceOne() {
	//1.使用内建方法make来创建一个slice,第一个参数为类型，第二个参数为长度，第三个参数为容量（不指定默认等于长度）
	//指定长度的话会填充类型的默认值，int是0，string是空字符串
	//sliceOne := make([]int, 5, 6)
	//sliceTwo := make([]int, 0, 6)
	//fmt.Println(len(sliceOne))
	//fmt.Println(len(sliceTwo))
	//fmt.Println(sliceOne)
	//fmt.Println(sliceTwo)
	//2.不允许创建长度大于容量的slice
	//sliceThree := make([]int, 3, 2)
	//fmt.Println(sliceThree)//会报错len larger than cap in make([]int)

	//3.slice字面量创建方式
	sliceFour := []int{1, 2, 3} //直接指定初始化的具体元素
	fmt.Println(sliceFour)
	//4.初始化索引
	sliceFive := []int{10: 1} //初始化一个有11个元素的切片
	fmt.Println(sliceFive)
	//5.创建一个nil的slice,做法是只声明不赋值，表示并不存在的slice
	var sliceNil []int
	fmt.Println(sliceNil)
	//6.创建一个empty的slice,做法是用过make函数初始化一个长度为0的slice即可，底层不会整整分配空间
	sliceEmpty := make([]int, 0)
	fmt.Println(sliceEmpty)
	//7.使用new创建，得到的是一个指针
}

func UseSlice() {
	//slice := make([]int, 5, 6)
	//1.slice通过索引正常赋值，若越界会报错panic: runtime error: index out of range
	//slice[0] = 12
	//slice[4] = 12
	//slice[5] = 12
	//fmt.Println(slice)
	//fmt.Println(slice)
	//2.通过旧的slice创建新的slice，左边是开始索引（包含），右边是结束索引（不包含），既然是索引页不能越界，同样会报错
	//特点：共享底层数组；通过这种方式创建的slice其实在底层操作的是同一块内存地址，也就是两个slice是引用关系，其中一个改变值，另一个对应的也会改变
	//newSlice := slice[2:5]
	//fmt.Println(slice)
	//fmt.Println(newSlice)
	//
	//slice[4] = 1212
	//fmt.Println(slice)
	//fmt.Println(newSlice)
	//3.容量只有被合并到长度的时候才可以访问
	//sliceOne := []int{10, 20, 30, 40}
	//4.append第一次增加容量是原来的2倍，不够的话是1.25倍
	//sliceTwo := make([]int,2,3)
	//sliceTwo = append(sliceTwo, 50)
	//fmt.Println(cap(sliceTwo))
	//sliceTwo = append(sliceTwo, 50)
	//fmt.Println(sliceTwo)
	//fmt.Println(cap(sliceTwo))
	// 附加一个新值到 slice，因为超出了容量，所以会创建新的底层数组,这样会断掉
	//newSlice := append(sliceOne, 50)
	//fmt.Println(sliceOne)
	//fmt.Println(newSlice)
	//sliceOne[0] = 1
	//fmt.Println(sliceOne)
	//fmt.Println(newSlice)
	//5.空切片和 nil 切片的区别在于，空切片指向的地址不是nil，指向的是一个内存地址，但是它没有分配任何内存空间，即底层元素包含0个元素。

}

func TestExtendArray() {
	//fmt.Println("实验一：数组共享内存地址")
	//sliceOne := make([]int, 1024, 1025)
	//sliceOne[4] = 12
	//fmt.Println("最开始sliceOne", sliceOne)
	//sliceTwo := sliceOne[3:5]
	//fmt.Printf("sliceOne地址 %p\n", sliceOne)
	//fmt.Printf("sliceTwo地址 %p\n", sliceTwo)
	//fmt.Println("从sliceOne创建的sliceTwo", sliceTwo)
	//sliceTwo[0] = 23
	//fmt.Println("将sliceTwo索引为0的位置的元素更改为12")
	//fmt.Println("更改后的sliceTwo", sliceTwo)
	//fmt.Println("sliceOne受影响变为", sliceOne)
	//fmt.Println("结论一：两个切片共享同一块内存地址")
	//fmt.Printf("结论二：创建的切片不是同一个，但底层数组内存地址是一个")
	//
	//
	//fmt.Println("实验二：测试扩容")
	//fmt.Printf("扩容前sliceOne地址 %p\n", sliceOne)
	//newSlice := append(sliceOne, 44, 66)
	//fmt.Printf("新的newSlice地址 %p\n", newSlice)
	//fmt.Printf("扩容后sliceOne地址 %p\n", sliceOne)
	//fmt.Printf("扩容后sliceTwo地址 %p\n", sliceTwo)
	//fmt.Println("结论三：扩容以后的切片长度小于等于容量的时候，两个切片底层数组是同一个地址，也就是说没有创建新数组，而是在原来旧的地址上紧挨着扩展了地址")
	//fmt.Println("扩容后的值：sliceOne", sliceOne)
	//fmt.Println("扩容后的容量：sliceOne", cap(sliceOne))
	//fmt.Println("扩容后的长度：sliceOne", len(sliceOne))
	//fmt.Println("新的值：newSlice", newSlice)
	//fmt.Println("新的容量：newSlice", cap(newSlice))
	//fmt.Println("新的长度：newSlice", len(newSlice))
	//fmt.Println("sliceTwo的值不受影响", sliceTwo)
	//fmt.Println("sliceTwo的容量", cap(sliceTwo))
	//fmt.Println("sliceTwo的长度", len(sliceTwo))
	//fmt.Println("###############################")
	//fmt.Println("结论四：当连续扩展2个以上元素的时候发现长度大于了原切片的容量的时候，newSlice指向了新的数组地址，并且容量是原来的2倍")
	//fmt.Println("###############################")
	//sliceOne[3] = 55
	//fmt.Println("新的值：newSlice", newSlice)
	//fmt.Println("新的值：sliceOne", sliceOne)
	//fmt.Println("新的值：sliceTwo", sliceTwo)
	//fmt.Println("结论五，当长度没有超过容量的时候，三者具有共享数组内存地址，更改一个，其他的也会受影响")
	//fmt.Println("sliceOne与sliceTwo始终会相互影响")

	//sliceOne = append(sliceOne, 44)
	//fmt.Printf("扩容后sliceOne地址 %p\n", sliceOne)
	//fmt.Printf("扩容后sliceTwo地址 %p\n", sliceTwo)
	//fmt.Println("结论三：扩容以后的切片长度小于等于容量的时候，两个切片底层数组是同一个地址，也就是说没有创建新数组，而是在原来旧的地址上紧挨着扩展了地址")
	//fmt.Println("扩容后的值：sliceOne", sliceOne)
	//fmt.Println("扩容后的容量：sliceOne", cap(sliceOne))
	//fmt.Println("扩容后的长度：sliceOne", len(sliceOne))
	//fmt.Println("sliceTwo的值不受影响", sliceTwo)
	//fmt.Println("sliceTwo的容量", cap(sliceTwo))
	//fmt.Println("sliceTwo的长度", len(sliceTwo))
	//fmt.Println("###############################")
	//fmt.Println("结论六：当连续扩展2个以上元素的时候发现长度大于了原切片的容量的时候，newSlice指向了新的数组地址，并且容量是原来的2倍")
	//fmt.Println("###############################")
	//sliceOne[3] = 55
	//fmt.Println("新的值：sliceOne", sliceOne)
	//fmt.Println("新的值：sliceTwo", sliceTwo)
	//fmt.Println("结论七，当在原数组基础上扩容后并且长度超过了容量，那么数组地址发生改变，这个时候sliceOne与sliceTwo无引用关系")
	//fmt.Println("结论八，当在原数组基础上扩容后并且长度没有超过容量，那么数组地址不改变，这个时候sliceOne与sliceTwo引用关系维持原状")

	//sliceOne := make([]int, 1024, 1024)
	//sliceOne = append(sliceOne, 44)
	//fmt.Println(cap(sliceOne))//1280/1024=1.25,证明1024个元素之后，是以不定倍数容量扩增的

	fmt.Println("切片拷贝copy函数：必须是2个已存在切片之间值的拷贝，当长度不一的时候只拷贝符合条件的数据，返回值是拷贝成功的元素")
	sliceOne := []int{1, 2, 3}
	newSlice := []int{0, 0}
	newLen := copy(newSlice, sliceOne)
	fmt.Println("newLen:", newLen)
	fmt.Println("sliceOne:", sliceOne)
	fmt.Println("newSlice:", newSlice)
	fmt.Printf("sliceOne:%p\n", sliceOne)
	fmt.Printf("newSlice:%p\n", newSlice)
}
