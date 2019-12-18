package main

import (
	"fmt"
	"log/mylogger"
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
	fileLogger := mylogger.NewFileLogger("info", "./", "log", 50)
	defer fileLogger.Close()
	for {
		fileLogger.Debug("这是一条Debug日志")
		fileLogger.TRACE("这是一条Trace日志")
		fileLogger.Info("这是一条Info日志")
		fileLogger.Warning("这是一条Warning日志")
		fileLogger.Error("这是一条Error日志")
		fileLogger.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second)
	}
}
