package main

import (
	"fmt"
)

func main() {

	//从控制台接受不了用户的信息
	//方式1  fmt.Scanln
	//先声明需要的变量
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试")
	fmt.Scanln(&isPass)
}
