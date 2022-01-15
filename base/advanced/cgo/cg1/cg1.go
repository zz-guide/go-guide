package main

/*
#include <stdio.h>

void sayHi() {
    printf("你好，我是C语言");
}
*/
import "C"

func main() {
	C.sayHi()
}
