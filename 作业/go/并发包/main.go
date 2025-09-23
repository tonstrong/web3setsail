package main

import (
	"fmt"
	"sync"
)

var instance int
var once sync.Once

func main() {
	//sync.Map
	//适用于读多写少场景，存在短暂的数据视图不一致，以换取更高的并发读性能。所以不适用于强一致性。
	var syncMap sync.Map
	syncMap.Store("a", "test")
	value, ok := syncMap.Load("a")
	if ok {
		fmt.Println("a=", value)
	}

	//syncMap的range
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println("key=", key, "value=", value)
		return true
	})

	//sync.Once
	//保证某个函数只会被执行一次，即使在多个 goroutine 同时调用的情况下。
	getInstance()

	//sync.Pool 是一个临时对象池，用于存储可以被多个 goroutine 复用的临时对象，以减少内存分配和垃圾回收压力。

}

func getInstance() int {
	once.Do(func() {
		instance = 1
	})
	return instance
}
