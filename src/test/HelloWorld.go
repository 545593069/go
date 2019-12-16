package test

import "fmt"

func main() {
	fmt.Println("Hello world")
	s := "aslkjahsdflkjhas"
	fmt.Println(s)
	s2 := `
		世情薄
		人情落
		雨送黄昏花易落
	`
	fmt.Println(s2)

	//多个判断条件
	//ifage>35{
	// fmt.Println("人到中年”)
	//}elseifage>18{
	// fmt.Println("青年")
	//}else{
	// fmt.Println(”好好学习! ")
	// }
	//作用域
	// age变量此时只在if条件判断语句中生效
	if age := 19; age > 18 { //如果age>18就执行这个
		fmt.Println("澳门首家线上赌场开业啦! ")
	} else { //否则就执行这个{}中的代码
		fmt.Println("改写暑假作业啦! ")
		fmt.Println(age)
	}
}
