package main

import "fmt"

//方法

//标识符:变量名函数名类型名方法名
// Go语言中如果标识符首字母是大写的，就表示对外部包可见(暴露的，公有的).

//Dog 这是一个狗的结构体
type dog struct {
	name string
}

//构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func (d dog) wang() {
	fmt.Printf("%s:汪汪汪···", d.name)
}

func main() {
	d1 := newDog("zhoulin")
	d1.wang()
}
