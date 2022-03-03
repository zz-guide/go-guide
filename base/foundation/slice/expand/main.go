package main

import (
	"log"
	"reflect"
	"unsafe"
)

/**
1.slice的源码在runtime.slice包中
2.append如果超出了原slice大小，则创建新的slice返回；否则返回原slice
*/

func main() {
	//TExpandRule()
	TSliceType()
}

/**
扩容规则：
	第一步：预估扩容后容量
	1.如果旧的容量乘以2 小于等于 所需最小容量，那么容量newCap就等于所需最小容量
	2.旧的容量乘以2 大于等于 所需最小容量。接下来判断长度len。
	3.如果旧的长度小于等于1024，那么直接在旧的容量基础上翻倍扩容；
	4.如果旧的长度大于等于1024，先在旧的基础上扩1.25倍，循环往复扩充

	1.oldCap * 2 < cap => newCap = cap
	   1.oldLen < 1024 => newCap = oldCap * 2
	   2.oldLen >= 1024 => newCap = oldCap *1.25(扩大1/4)

	第二步：分配合适规格内存
	所需内存 = 预估容量 * 元素类型大小
*/

func TExpandRule() {
	// 1.先看不用make的方式分配
	arr := []int{1, 2}
	// 不使用make分配，初始化几个值，cap=len=几
	log.Printf("扩容前: arr cap=%d, len=%d\n", cap(arr), len(arr))

	// 预估容量=5，5*8=40字节，最接近的是48字节，48/8=6，最后分配6
	arr = append(arr, 3, 4, 5)
	log.Printf("append 之后: arr cap=%d, len=%d\n", cap(arr), len(arr))

	ex := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// ex[:], ex[0:], ex[:0] 等同于浅拷贝自身
	newEx := ex[:]
	log.Printf("ex: %p %#v \n", &ex, ex)
	log.Printf("newEx: %p %#v \n", &newEx, newEx)
	log.Println("-------------------------")

	// 会同步修改原来的slice
	newEx[0] = 20
	log.Printf("ex: %p %#v \n", &ex, ex)
	log.Printf("newEx: %p %#v \n", &newEx, newEx)
	log.Println("-------------------------")

	// :左右必须是none negative即非负数，两边的数字表示0-capacity,且左边的必须小于右边的
	// 从index=left开始数，左开右必，含左不含右
	// [:2]等同于[0:2]
	newEx2 := ex[0:1]
	log.Printf("ex: %p %#v \n", &ex, ex)
	log.Printf("newEx2: %p %#v \n", &newEx2, newEx2)
	log.Println("-------------------------")

	// ex[0:0]，ex[1:1]，ex[2:2] 都表示重置切片，清空拥有的元素，数字不能大于capacity
	//newEx1 := ex[0:0]
	//log.Printf("ex: %p %#v \n", &ex, ex)
	//log.Printf("newEx1: %p %#v \n", &newEx1, newEx1)
	log.Println("-------------------------")

	// nil slice,nil 切片不能直接赋值使用，需要搭配append方法进行初始化，append底层判断容量不够就会进行初始化
	// 2个nil slice不能直接比较
	var a []int
	var b []int

	// empty slice, 所有空切片结构体Data指向的地址是同一个
	c := make([]int, 0)
	var d = make([]int, 0)
	var e []int
	var f []int

	log.Printf("a: %p %+v \n", &a, *(*reflect.SliceHeader)(unsafe.Pointer(&a)))
	log.Printf("b: %p %+v \n", &b, *(*reflect.SliceHeader)(unsafe.Pointer(&b)))
	log.Printf("c: %p %#v \n", &c, *(*reflect.SliceHeader)(unsafe.Pointer(&c)))
	log.Printf("d: %p %#v \n", &d, *(*reflect.SliceHeader)(unsafe.Pointer(&d)))
	log.Printf("e: %p %#v \n", &e, *(*reflect.SliceHeader)(unsafe.Pointer(&e)))
	log.Printf("f: %p %#v \n", &f, *(*reflect.SliceHeader)(unsafe.Pointer(&f)))
	log.Println("a nil slice", a == nil)
	log.Println("b nil slice", b == nil)
	log.Println("f empty slice", f == nil)

	// 检查slice是否为空，通过len()检测
	//log.Println("slice is empty:", len(a), cap(a), len(f), cap(f))

	// copy 深拷贝
	t1 := []int{1, 2, 3}
	t2 := make([]int, 3)
	copy(t2, t1)
	log.Printf("t1: %p %#v \n", &t1, t1)
	log.Printf("t2: %p %#v \n", &t2, t2)
	log.Println("-------------------------")

	t2[2] = 45
	log.Printf("t1: %p %#v \n", &t1, t1)
	log.Printf("t2: %p %#v \n", &t2, t2)
}

func F1() {
	// ptr,len,cap  ptr指向底层数组的开始，len表示长度
	// 结论：通过make创建的切片，初始值是类型的默认值，并且len=cap
	s := make([]int, 5)
	log.Println("s6,长度,容量:", s, len(s), cap(s), (*reflect.SliceHeader)(unsafe.Pointer(&s)))

	// 结论：通过make创建的切片，初始值是类型的默认值，并且len为5，cap为10
	// cap不能比len小，否则会编译失败，len larger than cap in make([]int)
	//s1 := make([]int, 3, 4)
	//log.Println("s1,长度,容量:", s1, len(s1), cap(s1))

	// 结论：如果append之后，还没有超出原数组的容量，那么，切片中的指针指向的位置，就还是原数组，如果append之后，超过了原数组的容量，那么，Go就会开辟一块新的内存，把原来的值拷贝过来，这种情况丝毫不会影响到原数组。
	a := s[:]
	log.Println("a,长度,容量:", a, len(a), cap(a), (*reflect.SliceHeader)(unsafe.Pointer(&a)))
	a = append(a, 1)
	// 结论：超出容量以后创建新的slice，不影响原来的slice
	log.Println("----扩容-----")
	log.Println("a,长度,容量:", a, len(a), cap(a), (*reflect.SliceHeader)(unsafe.Pointer(&a)))
	log.Println("s6,长度,容量:", s, len(s), cap(s), (*reflect.SliceHeader)(unsafe.Pointer(&s)))
}

func append1() {
	// 扩容规则定义再runtime.slice的growslice方法
	s1 := make([]int, 1022, 1022)
	log.Println("s1,长度,容量:", len(s1), cap(s1))
	// 结论：长度小于1024扩容为原来cap的2倍
	// 结论：长度大于1024扩容为原来cap的1.25倍
	s1 = append(s1, 1, 2)
	log.Println("s1,长度,容量:", len(s1), cap(s1))
}

func append2() {
	// 当len>cap的时候才会扩容
	s1 := make([]int, 1026, 1026)
	log.Println("s1,长度,容量:", len(s1), cap(s1))

	// 大于1024扩容1.25倍
	s1 = append(s1, 1)
	log.Println("s1,长度,容量:", len(s1), cap(s1))
	// 结论：扩容完之后还需要申请内存，内存对齐，所以不是准确的1.25倍，内存span都是4的倍数内存块
}

func F3() {
	//s6 := make([]struct{}, 0, 2)
	//log.Println("s6,长度,容量:", s6, len(s6), cap(s6))
	//a := struct{}{}
	//log.Println("a占多少个字节:", unsafe.Sizeof(a))
	//s6 = append(s6, a)
	//log.Println("s6,长度,容量:", s6, len(s6), cap(s6))

	var s []int
	s = nil
	log.Println("s6,长度,容量:", s, len(s), cap(s), s == nil, unsafe.Pointer(&s))
	s = append(s)
	log.Println("s6,长度,容量:", s, len(s), cap(s), s == nil, unsafe.Pointer(&s))
}

func TSliceType() {
	// nil切片和空切片的区别？
	var nilSlice []int // nil 切片,Data指向0，没有初始化

	// 所有空切片结构体Data指向的地址是同一个
	emptySlice1 := make([]int, 0) // 空切片
	emptySlice2 := make([]int, 0)

	log.Printf("nil slice pointer:%+v, emptySlice1 pointer:%+v, emptySlice2 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&nilSlice)), *(*reflect.SliceHeader)(unsafe.Pointer(&emptySlice1)), *(*reflect.SliceHeader)(unsafe.Pointer(&emptySlice2)))

	// nil slice 与 empty slice 不相等
	log.Printf("nil slice != empty slice %v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&nilSlice))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&emptySlice1))).Data)

	// empty slice 是相等的
	log.Printf("empty slice == empty slice %v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&emptySlice1))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&emptySlice2))).Data)

	// nil 切片不能直接赋值使用，需要搭配append方法进行初始化，append底层判断容量不够就会进行初始化

	// 直接panic
	//nilSlice[0] = 2
	nilSlice = append(nilSlice, 2)
	log.Println("nilSlice:", nilSlice)

	log.Println("越界panic nilSlice[5]:", nilSlice[5]) // index out of range [5] with length 1

}

func Pop() {
	a := []int{1, 2, 3, 4, 5}
	c := a[0:len(a)]
	// 0-长度减1可以把最后一个元素弹出
	b := a[0 : len(a)-1]

	log.Println("a:", a)
	log.Println("c:", c)
	log.Println("b:", b)
}

func references() {
	x := []int{1, 2, 3}

	y := x[:2]
	y = append(y, 50)
	log.Println("x:", x, ";y:", y)
	y = append(y, 60)
	log.Println("x:", x, ";y:", y)
}

func pointer() {
	arr := make([]int, 3, 5)
	arr[0] = 1
	arr[1] = 2
	//arr = append(arr, 1, 2)
	log.Printf("arr 长度:%d 容量:%d\n", len(arr), cap(arr))
	arr1 := _fff(arr)
	log.Printf("arr: %p, %v \n", arr, arr)
	log.Printf("arr: %p \n", &arr)
	log.Printf("arr1: %p  %v \n", arr1, arr1)
	log.Printf("arr1: %p \n", &arr1)
}

func _fff(nums []int) []int {
	//nums[0] = 5
	nums = append(nums, 4)
	log.Printf("nums: %p \n", nums)
	log.Printf("nums: %p \n", &nums)
	return nums
}

// Category slice定义方法，默认是指针类型的，不能写*
type Category []int

func (c Category) Add(n int) {
	for i, v := range c {
		c[i] = v + n
	}
}

func TestSliceOne() {
	//1.使用内建方法make来创建一个slice,第一个参数为类型，第二个参数为长度，第三个参数为容量（不指定默认等于长度）
	//指定长度的话会填充类型的默认值，int是0，string是空字符串
	//sliceOne := make([]int, 5, 6)
	//sliceTwo := make([]int, 0, 6)
	//log.Println(len(sliceOne))
	//log.Println(len(sliceTwo))
	//log.Println(sliceOne)
	//log.Println(sliceTwo)
	//2.不允许创建长度大于容量的slice
	//sliceThree := make([]int, 3, 2)
	//log.Println(sliceThree)//会报错len larger than cap in make([]int)

	//3.slice字面量创建方式
	sliceFour := []int{1, 2, 3} //直接指定初始化的具体元素
	log.Println(sliceFour)
	//4.初始化索引
	sliceFive := []int{10: 1} //初始化一个有11个元素的切片
	log.Println(sliceFive)
	//5.创建一个nil的slice,做法是只声明不赋值，表示并不存在的slice
	var sliceNil []int
	log.Println(sliceNil)
	//6.创建一个empty的slice,做法是用过make函数初始化一个长度为0的slice即可，底层不会整整分配空间
	sliceEmpty := make([]int, 0)
	log.Println(sliceEmpty)
	//7.使用new创建，得到的是一个指针
}

func UseSlice() {
	//slice := make([]int, 5, 6)
	//1.slice通过索引正常赋值，若越界会报错panic: runtime error: index out of range
	//slice[0] = 12
	//slice[4] = 12
	//slice[5] = 12
	//log.Println(slice)
	//log.Println(slice)
	//2.通过旧的slice创建新的slice，左边是开始索引（包含），右边是结束索引（不包含），既然是索引页不能越界，同样会报错
	//特点：共享底层数组；通过这种方式创建的slice其实在底层操作的是同一块内存地址，也就是两个slice是引用关系，其中一个改变值，另一个对应的也会改变
	//newSlice := slice[2:5]
	//log.Println(slice)
	//log.Println(newSlice)
	//
	//slice[4] = 1212
	//log.Println(slice)
	//log.Println(newSlice)
	//3.容量只有被合并到长度的时候才可以访问
	//sliceOne := []int{10, 20, 30, 40}
	//4.append第一次增加容量是原来的2倍，不够的话是1.25倍
	//sliceTwo := make([]int,2,3)
	//sliceTwo = append(sliceTwo, 50)
	//log.Println(cap(sliceTwo))
	//sliceTwo = append(sliceTwo, 50)
	//log.Println(sliceTwo)
	//log.Println(cap(sliceTwo))
	// 附加一个新值到 slice，因为超出了容量，所以会创建新的底层数组,这样会断掉
	//newSlice := append(sliceOne, 50)
	//log.Println(sliceOne)
	//log.Println(newSlice)
	//sliceOne[0] = 1
	//log.Println(sliceOne)
	//log.Println(newSlice)
	//5.空切片和 nil 切片的区别在于，空切片指向的地址不是nil，指向的是一个内存地址，但是它没有分配任何内存空间，即底层元素包含0个元素。

}

func TestExtendArray() {
	//log.Println("实验一：数组共享内存地址")
	//sliceOne := make([]int, 1024, 1025)
	//sliceOne[4] = 12
	//log.Println("最开始sliceOne", sliceOne)
	//sliceTwo := sliceOne[3:5]
	//log.Printf("sliceOne地址 %p\n", sliceOne)
	//log.Printf("sliceTwo地址 %p\n", sliceTwo)
	//log.Println("从sliceOne创建的sliceTwo", sliceTwo)
	//sliceTwo[0] = 23
	//log.Println("将sliceTwo索引为0的位置的元素更改为12")
	//log.Println("更改后的sliceTwo", sliceTwo)
	//log.Println("sliceOne受影响变为", sliceOne)
	//log.Println("结论一：两个切片共享同一块内存地址")
	//log.Printf("结论二：创建的切片不是同一个，但底层数组内存地址是一个")
	//
	//
	//log.Println("实验二：测试扩容")
	//log.Printf("扩容前sliceOne地址 %p\n", sliceOne)
	//newSlice := append(sliceOne, 44, 66)
	//log.Printf("新的newSlice地址 %p\n", newSlice)
	//log.Printf("扩容后sliceOne地址 %p\n", sliceOne)
	//log.Printf("扩容后sliceTwo地址 %p\n", sliceTwo)
	//log.Println("结论三：扩容以后的切片长度小于等于容量的时候，两个切片底层数组是同一个地址，也就是说没有创建新数组，而是在原来旧的地址上紧挨着扩展了地址")
	//log.Println("扩容后的值：sliceOne", sliceOne)
	//log.Println("扩容后的容量：sliceOne", cap(sliceOne))
	//log.Println("扩容后的长度：sliceOne", len(sliceOne))
	//log.Println("新的值：newSlice", newSlice)
	//log.Println("新的容量：newSlice", cap(newSlice))
	//log.Println("新的长度：newSlice", len(newSlice))
	//log.Println("sliceTwo的值不受影响", sliceTwo)
	//log.Println("sliceTwo的容量", cap(sliceTwo))
	//log.Println("sliceTwo的长度", len(sliceTwo))
	//log.Println("###############################")
	//log.Println("结论四：当连续扩展2个以上元素的时候发现长度大于了原切片的容量的时候，newSlice指向了新的数组地址，并且容量是原来的2倍")
	//log.Println("###############################")
	//sliceOne[3] = 55
	//log.Println("新的值：newSlice", newSlice)
	//log.Println("新的值：sliceOne", sliceOne)
	//log.Println("新的值：sliceTwo", sliceTwo)
	//log.Println("结论五，当长度没有超过容量的时候，三者具有共享数组内存地址，更改一个，其他的也会受影响")
	//log.Println("sliceOne与sliceTwo始终会相互影响")

	//sliceOne = append(sliceOne, 44)
	//log.Printf("扩容后sliceOne地址 %p\n", sliceOne)
	//log.Printf("扩容后sliceTwo地址 %p\n", sliceTwo)
	//log.Println("结论三：扩容以后的切片长度小于等于容量的时候，两个切片底层数组是同一个地址，也就是说没有创建新数组，而是在原来旧的地址上紧挨着扩展了地址")
	//log.Println("扩容后的值：sliceOne", sliceOne)
	//log.Println("扩容后的容量：sliceOne", cap(sliceOne))
	//log.Println("扩容后的长度：sliceOne", len(sliceOne))
	//log.Println("sliceTwo的值不受影响", sliceTwo)
	//log.Println("sliceTwo的容量", cap(sliceTwo))
	//log.Println("sliceTwo的长度", len(sliceTwo))
	//log.Println("###############################")
	//log.Println("结论六：当连续扩展2个以上元素的时候发现长度大于了原切片的容量的时候，newSlice指向了新的数组地址，并且容量是原来的2倍")
	//log.Println("###############################")
	//sliceOne[3] = 55
	//log.Println("新的值：sliceOne", sliceOne)
	//log.Println("新的值：sliceTwo", sliceTwo)
	//log.Println("结论七，当在原数组基础上扩容后并且长度超过了容量，那么数组地址发生改变，这个时候sliceOne与sliceTwo无引用关系")
	//log.Println("结论八，当在原数组基础上扩容后并且长度没有超过容量，那么数组地址不改变，这个时候sliceOne与sliceTwo引用关系维持原状")

	//sliceOne := make([]int, 1024, 1024)
	//sliceOne = append(sliceOne, 44)
	//log.Println(cap(sliceOne))//1280/1024=1.25,证明1024个元素之后，是以不定倍数容量扩增的

	log.Println("切片拷贝copy函数：必须是2个已存在切片之间值的拷贝，当长度不一的时候只拷贝符合条件的数据，返回值是拷贝成功的元素")
	sliceOne := []int{1, 2, 3}
	newSlice := []int{0, 0}
	newLen := copy(newSlice, sliceOne)
	log.Println("newLen:", newLen)
	log.Println("sliceOne:", sliceOne)
	log.Println("newSlice:", newSlice)
	log.Printf("sliceOne:%p\n", sliceOne)
	log.Printf("newSlice:%p\n", newSlice)
}
