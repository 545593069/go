package main

import "fmt"

//人们可能错误地期望看到a，b，c作为输出。您可能会看到的是c，c，c.这是因为循环的每次迭代都使用变量v的相同实例，
//所以每个闭包共享该单个变量。当闭包运行时，它会在执行fmt.Println时打印v的值，但是自goroutine启动以来v可能已被修改。
//要在启动时将v的当前值绑定到每个闭包，必须修改内部循环以在每次迭代时创建新变量。

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		v := v //更简单的是创建一个新的变量，使用一个看似奇怪的声明样式但在Go中运行正常：
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}
	for _ = range values {
		<-done
	}

	for _, v := range values {
		//一种方法是将变量作为参数传递给闭包：
		go func(u string) {
			fmt.Println(u)
			done <- true
		}(v)
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}
