package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//打开文件
func main() {
	//osRead()
	//readFromFilebyBufio()
	//ioUtil()
	writeFile()
}

func writeFile() {
	file, err := os.OpenFile("xxx.txt", os.O_APPEND|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		fmt.Printf("open file failed ,err %v", err)
		return
	}
	file.WriteString("加了一行\n")
	//使用os.O_APPEND模式无法读取已有文件内容
	var tmp [128]byte
	read, _ := file.Read(tmp[:])
	fmt.Println(string(tmp[:read]))
}

func osRead() {
	fileObj, err := os.Open("file.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	//记得关闭文件
	defer fileObj.Close()
	//读文件
	//var tmp = make([]byte,128)//指定读的长度
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n == 0 {
			return
		}
	}
}

// File Bio读取文件
func readFromFilebyBufio() {
	file0bj, err := os.Open("file.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	//记得关闭文件
	defer file0bj.Close()
	//创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(file0bj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed, err:%v", err)
			return
		}
		fmt.Print(line)
	}
}

// ioutil. ReadFile读取整个文件
func ioUtil() {
	content, err := ioutil.ReadFile("file.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}
