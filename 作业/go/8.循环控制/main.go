package main

import (
	"fmt"
	"time"
)

func main() {
	//声明方式1
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}

	//声明方式2
	var a int = 3
	for a > 2 {
		fmt.Println(a)
		a--
	}

	//声明方式3 无限循环
	for {
		fmt.Println("死循环")
		break
	}

	//方式4
	var arr [3]int = [3]int{1, 2, 3}
	slice := arr[0:2]

	var tmap map[string]string = make(map[string]string)
	tmap["1"] = "1"
	tmap["2"] = "2"
	for i, v := range slice {
		fmt.Println("slice[", i, "] = ", v)
	}
	for k, v := range tmap {
		fmt.Println(k, v)
	}

	var ch1 chan string
	ch1 = make(chan string)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			ch1 <- "hello"
		}
	}()

	for {
		select {
		case t := <-ch1:
			fmt.Println(t)
		case <-time.After(time.Second):
			fmt.Println("过了2秒")
			break
		}

	}
}
