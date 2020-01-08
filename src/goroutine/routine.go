package main

import (
	"fmt"
	"time"
)

// goroutine
func hello(i int) {
	fmt.Println("hello", i)
}

//程序启动之后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 100000; i++ {
		go hello(i) //开启一个单独的goroutine去执行hello函数(任务)
	}
	fmt.Println("main")
	time.Sleep(time.Second)
	// main函数结束了由main函数 启动的goroutine也都结束了
}
