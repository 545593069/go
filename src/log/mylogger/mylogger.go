package main

import "time"

func main() {
	for {
		logger := Logger{UNKNOWN}
		Logger.Debug(logger, "这是一条Debug日志")
		Logger.TRACE(logger, "这是一条Trace日志")
		Logger.Info(logger, "这是一条Info日志")
		Logger.Warning(logger, "这是一条Warning日志")
		Logger.Error(logger, "这是一条Error日志")
		Logger.Fatal(logger, "这是一条Fatal日志")
		time.Sleep(time.Second)
	}
}
