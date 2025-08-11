package main

import (
	"fmt"
	_ "goInitOrder/pkg"
	_ "goInitOrder/pkg1"
)

func main() {
	//go 初始化顺序测试
	fmt.Println("main invoke")
}
