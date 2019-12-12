package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//学生管理系统
//菜单函数
func showMenu() {
	fmt.Println("welcome sms !")
	fmt.Println(`
	1.查看所有学生
	2.添加学生
	3.修改学生
	4.删除学生
	5.退出
	`)
}

var smr studentMgr

func main() {
	smr = studentMgr{ //修改的全局的哪个变量
		allStudent: make(map[int64]student, 100),
	}
	for {
		//等待用户输入选项
		fmt.Println("请输入要执行的序号")
		var choice int
		fmt.Println(choice)
		for {
			showMenu()
			//等待用户输入选项
			fmt.Print("请输入序号:")
			var choice int
			fmt.Scanln(&choice)
			fmt.Println("你输入的是:", choice)
			// ?
			switch choice {
			case 1:
				smr.showStudents()
			case 2:
				smr.addStudent()
			case 3:
				smr.editStudent()
			case 4:
				smr.deleteStudent()
			case 5:
				os.Exit(1)
			default:
				fmt.Println("滚~")
			}
		}

		type point struct {
			X int `json:”zhoulin"`
			Y int `json:"baodelu"`
		}

		po1 := point{100, 200}
		//序列化
		b1, err := json.Marshal(po1)
		//如果出错了
		if err != nil {
			fmt.Printf("marshal failed, err:%v\n", err)
		}
		fmt.Println(string(b1))
		//反序列化:由字符串--> Go中的结构体变量
		str1 := `{"zhoulin":10010,"baodelu" :10086}`
		var po2 point //造一个结构体变量,准备接收反序列化的值
		err = json.Unmarshal([]byte(str1), &po2)
		if err != nil {
			fmt.Printf("unmarshal failed, err:%v\n", err)
		}
		fmt.Println(po2)

	}
}
