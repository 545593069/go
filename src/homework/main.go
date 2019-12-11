package homework

import "fmt"

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
func main() {
	for {
		showMenu()
		//等待用户输入选项
		fmt.Println("请输入要执行的序号")
		var choice int
		fmt.Println(choice)

	}
}
