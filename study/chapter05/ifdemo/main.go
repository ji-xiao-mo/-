package main

import (
	"fmt"
)

func main() {
	var age byte
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("你年龄大于18岁，要对自己的行为负责")
	} else {
		fmt.Println("年龄有些小，要学习良好的行为习惯")
	}
}
