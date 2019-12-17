package main

import "fmt"

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿") // defer把它 后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("呵呵呵") //一个函数中可以有多个defer语句
	defer fmt.Println("哈哈哈") //多个defer语言按照先进后出(后进先出)的顺序延迟执行(栈？)

	fmt.Println("end")
}
func main() {
	deferDemo()
}
