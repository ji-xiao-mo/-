package main

import (
	"fmt"
	"strings"
)

// 返回值是一个函数，形参无
// 累加器
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func makesuffix(suffix string) func(string) string {

	return func(name string) string {
		//如果没有后缀，则加上，负责返回原来名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {

	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16
	f2 := makesuffix(".suffix")
	fmt.Println(f2("winter"))
	fmt.Println(f2("winter"))

}
