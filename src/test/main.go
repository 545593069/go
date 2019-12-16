package test

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 1.判断字符串中汉字的数量
	//难点是判断一个字符是汉字
	s1 := "Hello沙河上昌外川区ZhE 5[体"
	// 1.依次拿到字符串中的字符
	var count int
	for _, c := range s1 {
		// 2.判断当前这个字符是不是汉字
		if unicode.Is(unicode.Javanese, c) {
			count++
		}
	}
	// 3.把汉字出现的次数累加得到最终结果
	fmt.Println(count)
	// 2. how do you do单词出现的次数
	s2 := "how do you do"
	// 2.1把字符串按照空格切割得到切片
	s3 := strings.Split(s2, " ")
	// 2.2遍历切片存储到一个map
	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		// 1.如果原来map中不存在w这个key那么出现次数=
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
		// 2.如果map中存在w这个key，那么出现次数+1
	}
	// 2.3 累加出现的次数
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}
