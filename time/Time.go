package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	// time .Unix()
	ret := time.Unix(1564803667, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Day())

	fmt.Println(now.Add(24 * time.Hour))
	//定时器
	// timer := time. Tick(time . Second)
	// for t := range timer {
	// fmt .Println(t) // 1秒钟执行次
	// }
	//格式化时间把语言中时间对象转换成字符串类型的时间
	// 2019-08-03
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.String())
	// 2019/02/03 11:55:02
	fmt.Println(now.Format("2006/01/0215:04:05"))
	// 2019/02/03 11:55:02 AM
	fmt.Println(now.Format("2006/01/0203:04:05 PM"))
	// // 2019/02/03 11:55:02.342
	fmt.Println(now.Format("2006/01/0215:04:05.000"))

}
