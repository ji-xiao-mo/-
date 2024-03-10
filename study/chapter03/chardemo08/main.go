package main

import (
	"fmt"
)

// 32位单精度
func main() {

	var c1 byte = 'a'

	fmt.Println(c1)
	//使用格式化输出%c
	fmt.Printf("c1= %c \n", c1)

	var c2 int = '韩'
	fmt.Println("c1= ", c2)
	fmt.Printf("c2= %c", c2)

}
