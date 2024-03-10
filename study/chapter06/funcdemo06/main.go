package main

import (
	"fmt"
)

var (
	//全局有效匿名函数
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	//在定义时调用
	//案例：其两个数的和，使用匿名函数的方式
	res1 := func(n1 int, n2 int) int {
		sum := n1 + n2
		return sum
	}(10, 32)
	fmt.Println(res1)

	//把匿名函数赋给a变量
	a := func(n1 int, n2 int) int {
		sum := n1 + n2
		return sum
	}
	fmt.Print(a(10, 323))

}
