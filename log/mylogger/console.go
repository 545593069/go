package mylogger

import (
	"errors"
	"fmt"
	"os"
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

// LoggerFile 结构体
type LoggerFile struct {
	level       LogLevel
	filePath    string
	fileName    string
	maxFileSize int64
	file0bj     *os.File
	errFile0bj  *os.File
}

func (f *LoggerFile) initFile() (err error) {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}
	errFile0bj, err := os.OpenFile(fullFileName+" .err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}
	//日志文件都已经打开了
	f.file0bj = fileObj
	f.errFile0bj = errFile0bj
	return nil
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, filePath string, fileName string, maxSize int64) *LoggerFile {
	logLevel, err := parsLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f := &LoggerFile{
		level:       logLevel,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxSize,
	}
	err = f.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return f

}

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

func (f *LoggerFile) enable(loglevel LogLevel) bool {
	return f.level <= loglevel
}

func (f *LoggerFile) log(level LogLevel, format string, a ...interface{}) {
	if f.enable(level) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, LineNo := getInfo(3)
		fmt.Fprintf(f.file0bj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05 "), getLogName(level), fileName, funcName, LineNo, msg)
		if level >= ERROR {
			fmt.Fprintf(f.errFile0bj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05 "), getLogName(level), fileName, funcName, LineNo, msg)
		}
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

func (f *LoggerFile) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}
func (f *LoggerFile) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}
func (f *LoggerFile) TRACE(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}
func (f *LoggerFile) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}
func (f *LoggerFile) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}
func (f *LoggerFile) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

func (f *LoggerFile) Close() bool {
	err2 := f.errFile0bj.Close()
	err := f.file0bj.Close()
	if err2 != nil {
		return false
	}
	if err != nil {
		return false
	}
	return true
}
