package main

import (
	"fmt"
)

func main() {
	//16进制
	var a int8 = 0xf
	var a1 int8 = 0xf
	//8进制
	var b int8 = 0o17
	var b1 int8 = 0o17
	//2进制
	var c int8 = 0b1111
	var c1 int8 = 0b1111

	fmt.Println(a == a1)
	fmt.Println(b == b1)
	fmt.Println(c == c1)
	fmt.Println(a == b)
	fmt.Println(a == c)
	fmt.Println(b == c)

	//自动类型转换时，默认位双精度浮点数，float64
	var f1 float32 = 10.0
	var f2 = 10.0
	fmt.Println(f1 == float32(f2))

	//虚数 默认也是双精度虚数complex128
	var cx1 complex64 = 1 + 2i
	cx2 := 1 + 2i
	var cx3 = complex(1, 2)
	fmt.Println(cx2 == cx3)
	fmt.Println(complex128(cx1) == cx2)

	//byte 类型 为uint8内置别名 可以把byte和uint8视为同一类型
	//同时字符串可以直接被转换成byte[]切片
	var s1 string = "hello world!"
	var bytes []byte = []byte(s1)
	fmt.Println(bytes)
	//byte 也可以转换成string
	fmt.Println(string(bytes) == s1)

	//rune类型时int32内置别名，可以视为同int32。rune是一种特殊的整数类型
	//go语言中，一个rune表示为一个Unicode码点，一般情况下，一个码点可以视为一个Unicode字符，有些比较特殊的Unicode字符由多个码点组成
	var rs []rune = []rune(s1)
	var rs2 []rune = []rune("世")
	fmt.Println(rs)
	fmt.Println(rs2)

	//字符串两种声明方式 双引号，自己写换行；反引号，可以直接使用回车换行，会自动生成换行符，两字符串一致。
	var ss1 = "hello \n world!\n"
	var ss2 = `hello 
 world!
`
	fmt.Println(ss1 == ss2)

	//byte rune string其实可以互相转换。string长度和[]byte一样，英文字母占用一个byte，rune和byte一样，而中文是三个byte。而对于rune来说，只是一个unicode编码，所以有多少个字符长度就是多少
	var brs string = "hi，go语言"
	var brsbytes []byte = []byte(brs)
	var brsrunes []rune = []rune(brs)
	fmt.Println(len(brs))
	fmt.Println(brs)
	fmt.Println(len(brsbytes))
	fmt.Println(brsbytes)
	fmt.Println(len(brsrunes))
	fmt.Println(brsrunes)

	//go中字节使用utf-8编码
}
