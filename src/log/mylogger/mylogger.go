package main

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

func main() {
	f()
}

func f1() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

func f() {
	for {
		logger := Logger{INFO}
		Logger.Debug(logger, "这是一条Debug日志")
		Logger.TRACE(logger, "这是一条Trace日志")
		Logger.Info(logger, "这是一条Info日志")
		Logger.Warning(logger, "这是一条Warning日志")
		Logger.Error(logger, "这是一条Error日志")
		Logger.Fatal(logger, "这是一条Fatal日志")
		time.Sleep(time.Second)
	}
}
