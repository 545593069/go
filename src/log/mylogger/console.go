package main

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	FATAL   LogLevel = 1
	DEBUG   LogLevel = 2
	WARNING LogLevel = 3
	TRACE   LogLevel = 4
	INFO    LogLevel = 5
	ERROR   LogLevel = 6
)

//// NewLog 构造函数
//func NewLog(LeverStr string) Logger {
//	level, err := parsLogLevel(LeverStr)
//	if err != nil {
//		panic(err)
//	}
//	return Logger{level: level}
//}

//func parsLogLevel(s string) (LogLevel, error) {
//	s = strings.ToLower(s)
//	switch s {
//	case "debug":
//		return DEBUG, nil
//	case "trace":
//		return TRACE, nil
//	case "info":
//		return INFO, nil
//	case "warning":
//		return WARNING, nil
//	case "error":
//		return ERROR, nil
//	case "fatal":
//		return FATAL, nil
//	default:
//		err := errors.New("日志级别设置错误，设置级别为; DEBUG; TRACE; INFO; WARNING; ERROR; FATAL")
//		return UNKNOWN, err
//	}
//}

func getLogName(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Logger 结构体
type Logger struct {
	level LogLevel
}

func (l Logger) enable(loglevel LogLevel) bool {
	return l.level <= loglevel
}

func (l Logger) log(level LogLevel, msg string) {
	now := time.Now()
	funcName, fileName, LineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05 "), getLogName(level), fileName, funcName, LineNo, msg)
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		l.log(DEBUG, msg)
	}
}
func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		l.log(INFO, msg)
	}
}
func (l Logger) TRACE(msg string) {
	if l.enable(TRACE) {
		l.log(TRACE, msg)
	}
}
func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		l.log(WARNING, msg)
	}
}
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		l.log(ERROR, msg)
	}
}
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		l.log(FATAL, msg)
	}
}

func getInfo(n int) (funcName, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Printf("runtime .Caller() failed\n")
		return
	}
	funcName = strings.Split(runtime.FuncForPC(pc).Name(), ".")[1]
	fileName = path.Base(file)
	return
}
