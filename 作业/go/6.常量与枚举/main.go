package main

import "fmt"

// 常量是在编译期确定的，所以定义时必须赋值，并且不能使用方法的返回值作为常量。
// 定义后无法被修改。
// 标准定义
const a string = "a"

// 类型推导
const b = "b"

// 批量定义
const c, d, e = 1, "d", false
const f, g, h int = 1, 2, 3
const (
	i string = "i"
	j int    = 1
	k bool   = true
)

// go中枚举是使用常量定义 本质就是一系列常量
type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

func (g *Gender) isMale() bool {
	return *g == Male
}

// iota关键词
// 自动为常量赋值，起始值为 行数-1
// 如果不写值，则都为最上面定义的那个值
type ConnState int

const (
	Init        ConnState = 0
	StateNew    ConnState = -1
	StateActive           = iota
	StateIdle
	StateClosed
)

func main() {
	fmt.Println(StateNew, StateActive, StateIdle, StateClosed)
}
