package main

import (
	"fmt"
)

func main() {
	//len函数是按照字节的个数来算作长度，汉字占三个字节
	str := "hello韩"
	fmt.Println("str len=", len(str)) // 8

	str2 := "hello北京"
	//字符串遍历，处理有中文的问题 r := []rune(str2)
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符为=%c\n", r[i])
	}

}
