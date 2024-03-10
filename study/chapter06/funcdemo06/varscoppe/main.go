package main

import (
	"fmt"
)

// 全局变量
var age = 10
var Name = "jack~"

// 函数
func test() {
	//作用域只在test函数内部
	age := 20
	Name := "tom~"
	fmt.Println(age)
	fmt.Println(Name)

}

func main() {
	test()
	fmt.Println(age)
}
