package main

import (
	"fmt"
)

// 定义一个struct结构体，将cat的各个字段属性，放入cat结构体中
type cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func main() {
	//使用struct来完成案例

	//创建一个cat变量
	var cat1 cat
	cat1.Name = "小白"
	cat1.Age = 10
	cat1.Color = "白"
	cat1.Hobby = "吃🐟"
	fmt.Println(cat1)
}
