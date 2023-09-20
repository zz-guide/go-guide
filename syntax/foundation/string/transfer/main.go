package main

import (
	"log"
	"reflect"
	"strconv"
	"unsafe"
)

func main() {
	//strToNumber()
	strToByte()
}

func strToNumber() {
	/**
	字符串转数字
	*/
	str := "123"
	log.Printf("str: %s %+v\n", str, &str)
	n, _ := stringToInt(str)
	log.Printf("n: %d %+v\n", n, &n)

}

func strToByte() {
	str := "hello world 你好"
	b1 := stringToByte1(str)
	log.Printf("b1: %s %p\n", b1, &b1)

	b2 := stringToByte2(str)
	log.Printf("b2: %s %p\n", b2, &b2)

	b3 := stringToByte3(str)
	log.Printf("b3: %s %p\n", b3, &b3)

	b4 := stringToByte4(str)
	log.Printf("b4: %s %p\n", b4, &b4)

	log.Printf("b4 == b3: %t \n", &b4 == &b3)
}

func stringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func stringToUin64(str string) (uint64, error) {
	res, err := stringToInt(str)
	if err == nil {
		return uint64(res), nil
	}

	return 0, err
}

func stringToUin32(str string) (uint32, error) {
	res, err := stringToInt(str)
	if err == nil {
		return uint32(res), nil
	}

	return 0, err
}

func stringToUin16(str string) (uint16, error) {
	res, err := stringToInt(str)
	if err == nil {
		return uint16(res), nil
	}

	return 0, err
}

func stringToUin8(str string) (uint8, error) {
	res, err := stringToInt(str)
	if err == nil {
		return uint8(res), nil
	}

	return 0, err
}

func stringToUint(str string) (uint, error) {
	res, err := stringToInt(str)
	if err == nil {
		return uint(res), nil
	}

	return 0, err
}

func stringToIn64(str string) (int64, error) {
	res, err := stringToInt(str)
	if err == nil {
		return int64(res), nil
	}

	return 0, err
}

func stringToIn32(str string) (int32, error) {
	res, err := stringToInt(str)
	if err == nil {
		return int32(res), nil
	}

	return 0, err
}

func stringToIn16(str string) (int16, error) {
	res, err := stringToInt(str)
	if err == nil {
		return int16(res), nil
	}

	return 0, err
}

func stringToInt8(str string) (int8, error) {
	res, err := stringToInt(str)
	if err == nil {
		return int8(res), nil
	}

	return 0, err
}

func stringToFloat64(str string) (float64, error) {
	res, err := stringToInt(str)
	if err == nil {
		return float64(res), nil
	}

	return 0, err
}

func stringToFloat32(str string) (float32, error) {
	res, err := stringToInt(str)
	if err == nil {
		return float32(res), nil
	}

	return 0, err
}

func stringToRune(str string) []rune {
	return []rune(str)
}

func stringToByte1(str string) []byte {
	/**
	适用于小字符串
	*/
	return []byte(str)
}

func stringToByte2(str string) []byte {
	/**
	适用于中等字符串
	*/
	b := make([]byte, len(str))
	copy(b, str)
	return b
}

func stringToByte3(s string) []byte {
	/**
	适用于超大字符串
	缺点：可能会导致内存错误，如越界、重复释放、空指针等问题。
	*/
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func stringToByte4(str string) []byte {
	/**
	适用于超大字符串
	避免SliceHeader，StringHeader出错
	*/
	return unsafe.Slice(unsafe.StringData(str), len(str))
}
