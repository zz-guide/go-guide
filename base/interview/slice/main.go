package main

func main() {

	/**
		slice

	1.slice由三个部分组成：
	data：元素存哪里
	len：存了多少个元素
	cap：可以存多少个元素


	2.make
	如果通过make的方式定义这个变量，不仅会分配这三部分结构，还会开辟一段内存作为它的底层数组。
	var ints []int = make([]int, 2, 5)
	这里make会为ints开辟一段容纳5个整型元素的内存,还会把它们初始化为整型的默认值 0。
	已经存储的元素是可以安全读写的，但是超出这个范围就属于越界访问，会发生panic。


	3.new
	这次我们看看字符串类型的slice，但是不用make，来试试new。
	 ps := new([]string)
	new一个slice变量同样会分配这三部分结构，但它不负责底层数组的分配，所以data=nil，len和cap都是0。new的返回值就是slice结构的起始地址，所以ps它就是个地址。
	此时这个slice变量还没有底层数组，像下面这样的操作是不允许的：
	(*ps)[0] = "eggo"
	那谁来给它分配底层数组呢？
	答案是：append
	*ps = append(*ps, "eggo")
	通过append的方式添加元素，append就会给它开辟底层数组。如下图所示，这里开辟了一个字符串元素的数组。







































































































	*/
}
