package main

import (
	"errors"
	"log"
	"reflect"
)

/**
内置类型 src/builtin/builtin.go
*/
func main() {
	TVariableTypes()
}

func TVariableTypes() {
	// int int8 int16 int32 int64
	var intV int = 1
	log.Printf("intV, 值=%d,类型=%s\n", intV, reflect.TypeOf(intV))

	// uint uint8 uint16 uint32 uint64
	var uintV uint8 = 1
	log.Printf("uintV, 值=%d,类型=%s\n", uintV, reflect.TypeOf(uintV))

	// byte type byte = uint8
	var byteV byte = 1
	log.Printf("byteV, 值=%d,类型=%s\n", byteV, reflect.TypeOf(byteV))

	// rune type rune = int32
	var runeV rune = 's'
	log.Printf("runeV, 值=%d,类型=%s\n", runeV, reflect.TypeOf(runeV))

	// bool
	var boolV bool = false
	log.Printf("boolV, 值=%t,类型=%s\n", boolV, reflect.TypeOf(boolV))

	// float32 float64
	var float32V float32 = 1.1
	log.Printf("float32V, 值=%f,类型=%s\n", float32V, reflect.TypeOf(float32V))

	// complex64 complex128
	var complex64V complex64 = 1.1
	log.Printf("complex64V, 值=%f,类型=%s\n", complex64V, reflect.TypeOf(complex64V))

	// string
	var stringV string = "ss"
	log.Printf("stringV, 值=%s,类型=%s\n", stringV, reflect.TypeOf(stringV))

	// uintptr
	var uintptrV uintptr = 1
	log.Printf("uintptrV, 值=%d,类型=%s\n", uintptrV, reflect.TypeOf(uintptrV))

	type S struct{}
	// any type any = interface{}
	var anyV any = S{}
	log.Printf("anyV, 值=%d,类型=%s\n", anyV, reflect.TypeOf(anyV))

	var interfaceV interface{}
	log.Printf("interfaceV, 值=%v,类型=%T\n", interfaceV, reflect.TypeOf(interfaceV))

	// error
	var errorV error = errors.New("出错了")
	log.Printf("errorV, 值=%d,类型=%s\n", errorV, reflect.TypeOf(errorV)) // errors.errorString

	// func
	var funcV = func() {}
	log.Printf("funcV, 值=%p,类型=%T\n", funcV, reflect.TypeOf(funcV))

	// array slice
	var arrayV [5]int
	log.Printf("arrayV, 值=%+v,类型=%s\n", arrayV, reflect.TypeOf(arrayV))
	var sliceV []int
	log.Printf("sliceV, 值=%+v,类型=%s\n", sliceV, reflect.TypeOf(sliceV))

	// type comparable interface{ comparable }
	// 此类型只能当做参数类型，不能当做变量类型，源码中有说明

	// 还有一些用作文档的类型 Type Type1  IntegerType FloatType ComplexType 等
}
