package main

import "fmt"

func main() {
	//go的变量使用方式1
	//指明变量类型，声明后·若不赋值，则使用默认值
	//初始化i
	var i int
	fmt.Println("i=", i)

	//第二种，赋值，不给类型，根据值来推导出类型
	var num = "casc"
	fmt.Println("num=", num)

}
