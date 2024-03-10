package main

import (
	"fmt"
)

func main() {
	//map的声明和注意事项
	var a map[string]string
	//使用map之前要先make，给map分配数据空间
	a = make(map[string]string, 10)
	a["姓名1"] = "韩宗铮"
	a["姓名2"] = "武松"
	fmt.Println(a)
}
