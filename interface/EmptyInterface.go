package main

import "fmt"

//类型断言1
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Printf("不对\n")
	} else {
		fmt.Printf("传进来是一个字符串：%s\n", str)
	}
}

// 类型断言2
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串:", t)
	case int:
		fmt.Println("是一个int:", t)
	case int64:
		fmt.Println("是一个int64:", t)
	case bool:
		fmt.Println("是一个bool:", t)
	case []int:
		fmt.Println("是一个slice:", t)
	case map[string]int:
		fmt.Println("是一个map[string]int:", t)
	case func():
		fmt.Println("是一个函数类型:", t)
	}
}
func f() {}
func main() {
	assign(1)
	// assign(100)
	assign2(true)
	assign2("哈哈哈")
	assign2(int64(200))
	assign2([]int{1, 2, 3})
	assign2(map[string]int{"a": 1})
	assign2(f)
}
