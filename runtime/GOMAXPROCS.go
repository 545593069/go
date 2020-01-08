package main

import (
	"fmt"
	"runtime"
	"time"
)

//Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数
//一个操作系统线程对应用户态多个goroutine。
//go程序可以同时使用多个操作系统线程。
//goroutine和OS线程是多对多的关系，即m:n
func a() {
	for i := 1; i < 100; i++ {
		fmt.Println("A", i)
	}
}
func b() {
	for i := 1; i < 100; i++ {
		fmt.Println("B", i)
	}
}

//单核跑
func oneMain() {
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}

//2核跑
func twoMain() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
func main() {
	//oneMain()
	twoMain()
}
