package main

import (
	"fmt"
)

type cat struct{}
type dog struct{}
type person struct{}

//定义的接口，直接用一个函数调用即可
type speaker interface {
	speak() //挨打了就要叫
}

func (C cat) speak() {
	fmt.Println("喵喵喵~")
}
func (d dog) speak() {
	fmt.Println("旺旺旺~")
}
func (P person) speak() {
	fmt.Println("哭了~")
}
func da(x speaker) {
	x.speak()
}
func main() {
	var c1 cat
	var d1 dog
	var p1 person
	da(c1)
	da(d1)
	da(p1)
}
