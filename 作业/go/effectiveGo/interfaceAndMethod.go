package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	// 未使用的导入和变量
	_ "time"
)

// 未使用的导入和变量
var _ = io.EOF
var _ = log.Default()

//// Handler 接口：只要有 ServeHTTP 方法，就是 Handler
//type Handler interface {
//    ServeHTTP(ResponseWriter, *Request)
//}

// 结构体作为Handler
type Counter struct {
	n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctr.n++
	fmt.Fprintln(w, "Counter = %d\n\n", ctr.n)
}

// 简化的结构体
type Counter2 int

func (ctr *Counter2) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintln(w, "Counter2 = %d\n\n", *ctr)
}

// 通道作为handler
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ch <- req
	fmt.Fprintln(w, "Chan = %#v\n\n", ch)
}

// 方法作为Handler
func ArgServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, os.Args)
}

// 接口检查
type Sayer interface {
	Say() string
}

type Dog struct{}
type Cat struct{}

func LetSay(s Sayer) {
	fmt.Println(s.Say())
}

func (d Dog) Say() string {
	return ("wangwang")
}

func main() {
	ctr := new(Counter)
	http.Handle("/counter", ctr)

	ctr2 := new(Counter2)
	http.Handle("/counter2", ctr2)

	ch := new(Chan)
	http.Handle("/ch", ch)

	http.Handle("/func", http.HandlerFunc(ArgServer))

	//静态检查有无实现接口 cat未实现 报错
	d := Dog{}
	LetSay(d)
	c := Cat{}
	LetSay(c)

	//动态检查最典型的就是 encoding/json 包。
	//json 包有个 Marshaler 接口，定义了 “如何把自己转换成 JSON”：
	var id interface{} = d
	m, ok := id.(json.Marshaler)
	if ok {
		m.MarshalJSON()
	}

	//编译时强制检查 Marshaler 接口
	var _ json.Marshaler = (*Dog)(nil)

}
