package main

import "fmt"

//不管什么牌子的车，都能跑
type car interface {
	drive()
}
type dazhong struct{}
type baoma struct{}
type xiandai struct{}

func (d dazhong) drive() {
	fmt.Println("大众车在行驶")
}
func (b baoma) drive() {
	fmt.Println("宝马车在行驶")
}
func (x xiandai) drive() {
	fmt.Println("现代车在行驶")
}
func driver(car car) {
	car.drive()
}
func main() {
	var d1 dazhong
	var b1 baoma
	var x1 xiandai
	driver(d1)
	driver(b1)
	driver(x1)
}
