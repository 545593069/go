package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup //引入sync.WaitGroup的方法

func goroute(i int) {
	defer wg.Done() //计数减一
	fmt.Println("hello", i)
	time.Sleep(time.Second)
}

func main() {
	wg.Add(2) //指定计数为2
	go func() {
		defer wg.Done()
		goroute(1)
		fmt.Println("匿名函数")
	}()
	wg.Wait() //等待计数为0，表示goroutine的线程都执行完成
	fmt.Println("main hello")
}
