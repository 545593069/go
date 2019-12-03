package main

import (
	"fmt"
	"os"
)

var (
	allStudnet map[int64]*student //变量声明
)

func deleteStudent() {
	// 1.请用户输入要删除的学生的序号
	var deleteID int64
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&deleteID)
	// 2.去allStudent这 个map中根据学号删除对应的键值对
	delete(allStudnet, deleteID)

}

func addStudent() {
	//向allStudent中添加一个 新的学生
	// 1.创建一 个新学生I
	// 2.追加到allStudent这个map中
	//向a1lStudent中添加一个新的学生
	// 1.创建一个新学生
	// 1.1获取用户输入
	var(
		id   int64
		name string
	)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&name)
	// 1.2造学生(调用student的构造函数)
	newStu := &student{}
	newStu.SetId(id)
	newStu.SetName(name)
	// 2.追加到allStudent这个map中
	allStudnet[id] = newStu
}

func showAllStudent() {
	//把所有的学生都打印出来
	for i:=0 ;i < len(allStudnet);i++{
		s := allStudnet[int64(i)]
		fmt.Printf("学号:%d姓名:%s\n", s.id, s.name)
	}
	for k, v := range allStudnet {
		fmt.Printf("学号:%d姓名:%s\n", k, v.name)
	}
}

func main() {
	allStudnet = make(map[int64]*student, 48) //初始化(开辟内存空间)
	for {
		// 1.打印菜单
		fmt.Println("欢迎光临学生管理系统!")
		fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
		4.退出
	`)
		fmt.Print("请输入你要干啥:")
		// 2.等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)
		// 3.执行对应的函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("滚~")

		}
	}
}
