package main

import "fmt"

//go语言中的全局变量、方法、结构体等，只有公开和非公开两种状态。
//非公开只能同一个package使用，公开则可以被其他package使用。
//定义公开与非公开通过命名的首字母大小写控制，公开的为首字母大写，非公开为小写。

type Other struct{}

type Person struct {
	// ``里面为字段标记，不会有任何影响，只有在反射时能读取到里面的内容
	Name string `json:"name" gorm:"column:name"`
	Age  int
	Call func(int) bool
	Map  map[string]string
	Ch   chan string
	Arr  []uint8
	//这是一个切片容器，专门用来存储任意类型的数据（因为 interface{} 是 Go 中的空接口，可表示任何类型）
	//切片中的元素可以是不同类型（例如同时存放 int、string、struct 等）
	//本质上是对 interface{} 类型元素的动态长度序列
	Slice []interface{}
	Ptr   *int
	O     Other
}

// 匿名字段，相同类型只能存在一个。使用时通过类型直接访问 custom.string
type Custom struct {
	int
	string
	Other string
}

// 匿名结构体：没有定义名称的结构体。通常用于单元测试接口的参数
var test = struct {
	var1 string
	var2 int
}{} //最后的{}代表初始化，必须立即初始化

// 上面完整的写法
var test2 = struct {
	var1 string `a:"hh" b:"hhh"`
	var2 string `a:"haha" b:"hahaha"`
}{
	var1: "hello",
	var2: "world",
}

// 在函数方法内部声明匿名结构体
func method() {
	test := struct {
		var1 string `a:"hh" b:"hhh"`
		var2 string `a:"haha" b:"hahaha"`
	}{
		var1: "hello",
		var2: "world",
	}
	fmt.Println(test)
}

// 嵌套结构体
type A struct {
	a string
}

type B struct {
	A
	b string
}

type C struct {
	A
	B
	a string
	b string
}

// 定义结构体方法，相较于普通函数多了一个接收者 func后有个小括号引用变量名 和 结构体类型
// 结构体方法没有重写规则，即不允许相同名称方法存在，即使参数不同。
// 结构体方法可以定义在同一包中不同文件下，一般情况不这么做，除非特别复杂需要拆分。
func (a A) setAa(p string) {
	a.a = p
}

func (a A) getAa() string {
	return a.a
}

// 还可以定义接收者为指针类型的方法，两者可以用混用。但是如果要修改字段，通过接收者为值的方法是修改不了的，必须通过接收者为指针的方法。
// 所以将一个变量赋值给一个新的变量，新的变量修改并不影响老的，相当于是两个变量了。
func (a *A) getPAa() string {
	return a.a
}
func (a *A) setPAa(p string) {
	a.a = p
}

func main() {
	//访问嵌套结构体同名的变量，必须指定结构体，否则访问的是本级的变量
	a := A{a: "hello"}
	b := B{A: a, b: "world"}
	c := C{A: a, B: b, a: "chello", b: "cworld"}
	fmt.Println(c.a)
	fmt.Println(c.A.a)
	fmt.Println(c.B.A.a)
	fmt.Println(c.B.b)

	fmt.Println("=========================")

	//调用接收者为值 或 指针的结构体方法都可以
	fmt.Println(a.getAa())
	fmt.Println(a.getPAa())

	pa := &a
	fmt.Println(pa.getAa())
	fmt.Println(pa.getPAa())

	//如果要修改字段，通过接收者为值的方法是修改不了的，必须通过接收者为指针的方法。
	a.setAa("haha")
	fmt.Println(a.getAa())
	fmt.Println(a.getPAa())

	a.setPAa("haha")
	fmt.Println(a.getAa())
	fmt.Println(a.getPAa())

	pa.setAa("haha")
	fmt.Println(pa.getAa())
	fmt.Println(pa.getPAa())

	pa.setPAa("haha")
	fmt.Println(pa.getAa())
	fmt.Println(pa.getPAa())

	//所以将一个变量赋值给一个新的变量，新的变量修改并不影响老的，相当于是两个变量了。
	copy := a
	copy.setPAa("hahaha")
	fmt.Println(copy.getAa())
	fmt.Println(a.getAa())
	//必须使用指针才可以同步修改，不然就是副本，是深拷贝
	copya := &a
	(*copya).a = "hellocopy"
	fmt.Println(a.getAa())

}
