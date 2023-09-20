package main

import (
	"fmt"
	"log"
	"time"
)

/**
结论：
1.
for range通过下标遍历元素的性能跟for相差不大
for range直接遍历元素的性能比for慢近1000倍
for range通过下标遍历元素的性能跟for相差不大
for range直接遍历元素的性能在元素为小对象的情况下跟for相差不大，在元素为大对象的情况下比for慢很多
2.支持数组，切片，map,channel,string等迭代
3.循环周期内会重复使用同一个变量
4.允许返回单个值，从左到右代表索引，值
5.允许 _ 忽略返回值

*/
func main() {

	//TForRangeString()
	//TForRangeArray()
	//TForRangeSlice()
	//TForRangeChannel()
	//TForRangeMap()
	TForRangeCURD()
}

func TForRangeCURD() {
	// 参考：https://www.jb51.net/article/262025.htm
	arr := []interface{}{[]int{1}, []int{2}, []int{3}}
	log.Printf("arr地址=%p\n", &arr)
	for i, v := range arr {
		// 结论：range产生元素的副本，且循环周期内使用的是同一个变量，每次进行覆盖值，修改副本不会对原先的值产生影响
		// 只产生一个副本变量，循环使用
		v = []int{7}
		// 结论：中途添加元素会改变arr，但是不会改变循环次数
		// arr = append(arr, 8)
		// log.Printf("新arr1=%+v，地址=%p\n", arr[5], &arr)
		//if len(arr) > 0 {
		//	arr = arr[:len(arr)-1]
		//}

		log.Printf("新arr1=%+v，地址=%p\n", arr, &arr)
		log.Printf("索引：i=%d,值=%+v\n", i, v)

	}

	log.Println("新arr=", arr)
}

func TForRangeChannel() {
	myChannel := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			myChannel <- i
		}
	}()

	go func() {
		for c := range myChannel {
			fmt.Printf("value:%d\n", c)
		}
	}()
}

func TForRangeMap() {
	// 遍历map时没有指定循环次数，循环体与遍历slice类似。由于map底层实现与slice不同，map底层使用hash表实现，插入数据位置是随机的，所以遍历过程中新插入的数据不能保证遍历到。
	myMap := map[int]string{1: "语文", 2: "数学", 3: "英语"}
	for key, value := range myMap {
		fmt.Printf("key:%d,value:%s\n", key, value)
		fmt.Printf("key:%d,value:%s\n", key, myMap[key])
	}
}

func TForRangeArray() {
	myArray := [3]int{1, 2, 3}
	for i, ele := range myArray {
		fmt.Printf("index:%d,element:%d\n", i, ele)
		fmt.Printf("index:%d,element:%d\n", i, myArray[i])
	}
}

func TForRangeSlice() {
	mySlice := []string{"I", "am", "peachesTao"}
	for i, ele := range mySlice {
		fmt.Printf("index:%d,element:%s\n", i, ele)
		fmt.Printf("index:%d,element:%s\n", i, mySlice[i])
	}
}

func TForRangeString() {
	s := "peachesTao"
	// 结论：循环体中string中的元素实际上是byte类型，需要转换为字面字符
	for i, item := range s {
		fmt.Println(string(item))
		fmt.Printf("index:%d,element:%s\n", i, string(s[i]))
	}
}
