package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//在条件判断中声明的变量，若执行了判定条件，那么便会声明这个变量，后续可以使用
	var a int = 10
	if b := 1; a > 10 {
		b = 2
		fmt.Println("a > 10")
	} else if c := 3; b > 1 {
		b = 3
		fmt.Println("b > 1")
	} else {
		if c == 3 {
			fmt.Println("c == 3")
		}
	}

	//使用函数进行声明
	if a := method(); a {
		fmt.Println(a)
	}

	if b := throwError(); b != nil {
		fmt.Println(b)
	}

	//switch语句 基于不同条件的不同动作。
	//每个case分支都是唯一的，从上往下依次匹配，直到命中。如果某些case重复会报错
	//默认情况下自带break，无需在每个case声明break
	switch a := radom(); a {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("default")
	}

	var t interface{}
	t = &a
	switch v := t.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case float32:
		fmt.Println("float32", v)
	case *int:
		fmt.Println("*int", *v)
	default:
		fmt.Println("default", v)
	}
}

func method() bool {
	return true
}

func throwError() error {
	return errors.New("error!!!")
}

// 初始化随机种子
func init() {
	rand.Seed(time.Now().UnixNano())
}

func radom() int {
	return rand.Intn(10)
}
