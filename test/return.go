package main

// Go语言中函数的return不是原子操作，在底层是分为两步来执行
//第一步:返回值赋值
//第二步:真正的RET返回
//函数中如果存在defer,那么defer执行的时机是在第一步和第二步之间
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func main() {
	println(f1())
	println(f2())
	println(f3())
	println(f4())
}
