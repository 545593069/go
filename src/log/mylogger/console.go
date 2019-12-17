package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG   LogLevel = 1
	TRACE   LogLevel = 2
	INFO    LogLevel = 3
	WARNING LogLevel = 4
	ERROR   LogLevel = 5
	FATAL   LogLevel = 6
)

func parsLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("日志级别设置错误，设置级别为; DEBUG; TRACE; INFO; WARNING; ERROR; FATAL")
		return UNKNOWN, err
	}
}

// Logger 结构体
type Logger struct {
	level LogLevel
}

// NewLog 构造函数
func NewLog(LeverStr string) Logger {
	level, err := parsLogLevel(LeverStr)
	if err != nil {
		panic(err)
	}
	return Logger{level: level}
}

func (l Logger) enable(loglevel LogLevel) bool {
	return l.level <= loglevel
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		fmt.Printf("[%s] [DEBUG] %s\n", now.Format("2006-01-02 15:04:05 "), msg)
	}
}
func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		fmt.Printf("[%s] [INFO] %s\n", now.Format("2006-01-02 15:04:05 "), msg)
	}
}
func (l Logger) TRACE(msg string) {
	if l.enable(TRACE) {
		now := time.Now()
		fmt.Printf("[%s] [TRACE] %s\n", now.Format("2006-01-02 15:04:05 "), msg)
	}
}
func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		now := time.Now()
		fmt.Printf("[%s] [WARNING] %s\n", now.Format("2006-01-02 15:04:05 "), msg)
	}
}
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		fmt.Printf("[%s] [ERROR] %s\n", now.Format("2006-01-02 15:04:05 "), msg)
	}
}
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		now := time.Now()
		fmt.Printf("[%s] [FATAL] %s\n", now.Format("2006-01-02 15:04:05 "), msg)
	}
}
