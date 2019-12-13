package animal

import "fmt"

// 接口的实现
type animal interface {
	move()
	eat(string)
}
type cat struct {
	name string
	feet int8
}

func (C cat) move() {
	fmt.Println("走猫步..")
}
func (C cat) eat(x string) {
	fmt.Println("猫吃鱼...", x)
}

type chicken struct {
	feet int8
}

func (C chicken) move() {
	fmt.Println("走鸡步..")
}
func (C chicken) eat() {
	fmt.Println("鸡吃虫...")
}

func run(a animal) {
	a.move()
}

func eat(a animal) {
	a.eat("")
}
