package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//声明 初始化零值为nil
	var p1 *int
	var p2 *string
	//初始化指针时，必须通过另外一个变量：
	var a int = 100
	var b string = "hello"
	//或者使用一个声明变量赋值给一个指针
	// p1 := &<struct type>{}

	p1 = &a
	p2 = &b
	fmt.Println(p1, p2)

	//获取指针的指针，即指针的内存地址 一般是在反射或者比较底层的框架使用
	// var p2 **<type>
	var p3 **string = &p2
	fmt.Println(p3)

	//访问指针指向的值
	fmt.Println(*p1, *p2, *p3, **p3)

	//指针无法像C一样进行加减计算操作 从而进行偏移访问，但是可以使用其他方式。使用的时候需要小心！
	//unsafe.Point 和 uintptr
	//先将指针转换成 unsafe.Point，再转换成uintptr
	var up *string
	var ua string = "hello"
	up = &ua
	up1 := unsafe.Pointer(up)
	up2 := unsafe.Pointer(&ua)

	//unsafe.Pointer转uintptr
	uintptr1 := uintptr(up1)
	uintptr2 := uintptr(up2)
	fmt.Println(uintptr1)
	fmt.Println(uintptr2)

	//uintptr转uintptr地址
	uintptr1 = uintptr1 + 1
	//uintptr1 是一个表示内存地址的整数（uintptr 类型）
	//unsafe.Pointer(uintptr1) 将地址值转换为通用指针类型
	//(*uint8)(...) 再将通用指针强转为指向 uint8 类型的指针
	//最终 c 就是一个指向 uintptr1 所表示的内存地址、且认为该地址存储的是 uint8 类型数据的指针
	c := (*uint8)(unsafe.Pointer(uintptr1))
	fmt.Println(c)
	fmt.Println(*c)

}
