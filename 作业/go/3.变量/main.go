package main

import "fmt"

// 全局变量
// 自动推导类型
var var1 = "var1"

// 初始为空串
var var2 string

// 初始为false
var var3 bool

// 初始为0
var var4 int

// 数组 切片 等为nil
var var5 []int

// 统一集中声明
var (
	var6  string
	var7  int
	var8  map[string]int
	var9  []byte
	var10 *int
)

func main() {
	//局部变量
	var invar1 int = 1
	fmt.Println(invar1)
	fmt.Println(var6)
	fmt.Println(var7)
	fmt.Println(var8)
	fmt.Println(var9)
	fmt.Println(var10)

	//虽然打印出来map是空的，但是直接使用会报错
	//var8["1"] = 1
	var8 = make(map[string]int, 0)
	var8["1"] = 1
	fmt.Println(var8)

	var a, b = test()
	fmt.Println(a, b)
}

// 还可以在方法声明的时候在函数结果处声明变量，这种方式全局变量是没有的，同时也不是必须要给变量复制或者使用
func test() (a int, b int) {
	a = a + 1
	b = b + 2
	return 0, 0
}
